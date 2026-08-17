/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package inst

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/bytedance/mockey/internal/monkey/common"
	"golang.org/x/arch/x86/x86asm"
)

// branchLen is the width of the branch sequence Disassemble is asked to make
// room for on amd64: MOV RDX, imm64 (10 bytes) + JMP RDX (2 bytes).
// It matches len(BranchInto(...)), which is what patch.go passes as `required`.
const branchLen = 12

// ret is a single-byte RET (0xc3), i.e. the shortest possible function body.
const ret = 0xc3

// oracle builds a stand-in for withinTargetFunc that answers "the target
// function's slot is [0, size)". The real oracle reads pclntab through
// runtime.FuncForPC and therefore only works on a live function entry, which a
// synthetic byte slice is not; injecting it lets these tests state the function
// boundary exactly instead of hoping the compiler lays one out for them.
func oracle(size int) func([]byte, int) (bool, bool) {
	return func(_ []byte, off int) (bool, bool) { return off < size, true }
}

// unknownOracle is the "FuncForPC could not answer" case (pc outside the text
// segment, e.g. a heap slice). It must make Disassemble fall back to the padding
// rule rather than refuse.
func unknownOracle(_ []byte, _ int) (bool, bool) { return false, false }

// TestDisassembleSecondReturnPath is the point of this change: a function whose
// first RET is followed by *more of its own code* (a second return path) must be
// patchable, because patching the entry makes every one of its instructions
// unreachable. Before, any non-0xCC byte after the first RET was refused with
// "function is too short to patch", even though the function was not short.
//
// The three shapes are the ones measured in the field, keyed by where the first
// RET ends: 8, 9 and 10. All are nil-guarded getters of the
// (*BaseResp).GetStatusCode family, which is why the 8 and 9 cases dominate.
func TestDisassembleSecondReturnPath(t *testing.T) {
	tests := []struct {
		name     string
		code     []byte
		codeSize int // instruction bytes of the target, padding excluded
		slotSize int // target's pclntab slot: its code *plus* its padding
		retEnd   int
		want     int
	}{
		{
			// RET_end=8, the exact shape dumped from the field:
			//   +0 TEST RAX,RAX | +3 JE +3 | +5 MOV EAX,[RAX] | +7 RET
			//   +8 XOR EAX,EAX (nil branch) | +10 RET | +11 padding
			name: "ret ends at 8 then nil branch",
			code: []byte{
				0x48, 0x85, 0xc0, // TEST RAX, RAX
				0x74, 0x03, // JE +3
				0x8b, 0x00, // MOV EAX, [RAX]
				ret,        // RET            @7, ends at 8
				0x31, 0xc0, // XOR EAX, EAX   @8
				ret,     // RET               @10
				padByte, // padding starts at 11
				padByte, padByte, padByte, padByte,
			},
			codeSize: 11,
			slotSize: 16,
			retEnd:   8,
			want:     branchLen,
		},
		{
			// RET_end=9: one byte wider prologue, same story. The instruction that
			// straddles `required` here ends exactly on it.
			name: "ret ends at 9 then nil branch",
			code: []byte{
				0x48, 0x85, 0xff, // TEST RDI, RDI
				0x74, 0x04, // JE +4
				0x8b, 0x47, 0x08, // MOV EAX, [RDI+8]
				ret,        // RET           @8, ends at 9
				0x31, 0xc0, // XOR EAX, EAX  @9
				ret,     // RET              @11
				padByte, // padding at 12
				padByte, padByte, padByte,
			},
			codeSize: 12,
			slotSize: 16,
			retEnd:   9,
			want:     branchLen,
		},
		{
			// RET_end=10, and here the second return path crosses `required`: the
			// XOR at [10,12) ends at 12 == required, then RET at 12 is past it, so
			// the cutting index is exactly branchLen.
			name: "ret ends at 10 then nil branch",
			code: []byte{
				0x48, 0x85, 0xc0, // TEST RAX, RAX
				0x74, 0x05, // JE +5
				0x0f, 0xb6, 0x40, 0x10, // MOVZX EAX, [RAX+16]
				ret,        // RET           @9, ends at 10
				0x31, 0xc0, // XOR EAX, EAX  @10..11
				ret,     // RET              @12
				padByte, // padding at 13
				padByte, padByte,
			},
			codeSize: 13,
			slotSize: 16,
			retEnd:   10,
			want:     branchLen,
		},
		{
			// RET_end=11, measured on a getter that computes before returning
			// (r.A * 2). Included because a census of real compiled getters turned
			// up 8, 10 and 11, so 11 is a shape that actually occurs.
			name: "ret ends at 11 then nil branch",
			code: []byte{
				0x48, 0x85, 0xc0, // TEST RAX, RAX
				0x74, 0x06, // JE +6
				0x8b, 0x00, // MOV EAX, [RAX]
				0x01, 0xc0, // ADD EAX, EAX
				0x90,                         // NOP
				ret,                          // RET           @10, ends at 11
				0xb8, 0xff, 0xff, 0xff, 0xff, // MOV EAX, -1  @11..15
				ret, // RET @16
				padByte, padByte, padByte, padByte, padByte, padByte, padByte,
				padByte, padByte, padByte, padByte, padByte, padByte, padByte, padByte,
			},
			codeSize: 17,
			slotSize: 32,
			retEnd:   11,
			// The MOV at [11,16) straddles `required`, so the cutting index rounds
			// up to its end rather than stopping at 12.
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Guard the fixture itself: assert the first RET really ends where the
			// case name claims, so a typo in the bytes cannot silently turn this
			// into a test of some other shape.
			if got := firstRetEnd(t, tt.code); got != tt.retEnd {
				t.Fatalf("fixture first RET ends at %d, want %d", got, tt.retEnd)
			}
			if tt.retEnd >= branchLen {
				t.Fatalf("fixture is not a short-function case: RET ends at %d >= %d",
					tt.retEnd, branchLen)
			}

			if tt.slotSize != len(tt.code) {
				t.Fatalf("fixture slot %d should span the whole buffer %d",
					tt.slotSize, len(tt.code))
			}
			got := disassemble(tt.code, branchLen, true, oracle(tt.slotSize))
			if got != tt.want {
				t.Fatalf("cutting index = %d, want %d", got, tt.want)
			}
			if got < branchLen {
				t.Fatalf("cutting index %d is shorter than the branch sequence %d",
					got, branchLen)
			}
			if !onInstructionBoundary(t, tt.code, got) {
				t.Fatalf("cutting index %d is not an instruction boundary", got)
			}
			// The relaxation must not reach outside the target: everything we are
			// about to overwrite has to be the target's own code or its padding.
			if got > tt.codeSize {
				pad := tt.code[tt.codeSize:got]
				for i, b := range pad {
					if b != padByte {
						t.Fatalf("byte %d past the function end is %#x, not padding",
							tt.codeSize+i, b)
					}
				}
			}
		})
	}
}

// TestDisassembleRefusesNeighbour is the negative half, and it is the load-bearing
// one: proving the old failures now pass says nothing unless the check can still
// fail. Each case puts a real neighbouring function inside the window that the
// branch sequence would overwrite, and Disassemble must refuse every one.
func TestDisassembleRefusesNeighbour(t *testing.T) {
	// MOV RBP,RSP; SUB RSP,0x10; TEST RAX,RAX; JNE -52. The last two bytes are
	// deliberately "75 cc": a JNE whose displacement byte is 0xCC, which a naive
	// byte-wise padding scan would mistake for padding.
	neighbour := []byte{0x48, 0x89, 0xe5, 0x48, 0x83, 0xec, 0x10, 0x48, 0x85, 0xc0, 0x75, 0xcc}

	tests := []struct {
		name     string
		code     []byte
		slotSize int // where the target's slot ends and the neighbour begins
	}{
		{
			// The minimal function: RET and nothing else, next function abutting.
			name:     "one byte function then a neighbour",
			code:     append([]byte{ret}, neighbour...),
			slotSize: 1,
		},
		{
			// A second return path that ends before `required`, with the neighbour
			// starting immediately after: the relaxation may consume the target's
			// own tail but must stop at its boundary.
			name: "second return path then a neighbour",
			code: append([]byte{
				0x48, 0x85, 0xc0, // TEST RAX, RAX
				0x74, 0x03, // JE +3
				0x8b, 0x00, // MOV EAX, [RAX]
				ret,        // RET @7
				0x31, 0xc0, // XOR EAX, EAX
				ret, // RET @10, function ends at 11
			}, neighbour...),
			slotSize: 11,
		},
		{
			// Some padding, but not enough: the neighbour still begins inside the
			// window we would clobber.
			name:     "padding that runs out before the branch width",
			code:     append([]byte{0x48, 0x89, 0xc8, ret, padByte, padByte}, neighbour...),
			slotSize: 4,
		},
		{
			// The subtle one: the neighbour's first instruction *starts* inside the
			// target's slot only by one byte... here instead the target's last
			// instruction would have to extend past its own end. Boundary is 12, and
			// the 4-byte IMUL at [10,14) crosses it, so it must be refused even
			// though offset 10 itself is still inside.
			name: "instruction starting inside but ending outside",
			code: append([]byte{
				0x48, 0x85, 0xc0, // TEST RAX, RAX
				0x74, 0x02, // JE +2
				0x8b, 0x00, // MOV EAX, [RAX]
				0x90,                   // NOP
				0x90,                   // NOP
				ret,                    // RET @9
				0x48, 0x0f, 0xaf, 0xc0, // IMUL RAX, RAX @10..13, crosses the end
			}, neighbour...),
			slotSize: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := recoverOf(func() { disassemble(tt.code, branchLen, true, oracle(tt.slotSize)) })
			if err == nil {
				t.Fatal("expected Disassemble to refuse writing over the neighbour, but it returned")
			}
			// The refusal must be the length check, not some incidental panic
			// (a decode error, an index out of range) that would make this test
			// pass for the wrong reason.
			if msg := fmt.Sprint(err); !strings.Contains(msg, "too short to patch") {
				t.Fatalf("refused for an unexpected reason: %v", msg)
			}
		})
	}
}

// TestDisassembleUnknownBoundaryFallsBackToPadding pins the behaviour when the
// pclntab lookup cannot answer (nil FuncForPC, e.g. a pc outside the text
// segment). Then the old padding rule applies unchanged: padding is accepted,
// real code is refused. This is what keeps a failed lookup from turning a patch
// that used to work into a refusal.
func TestDisassembleUnknownBoundaryFallsBackToPadding(t *testing.T) {
	padded := append([]byte{0x48, 0x89, 0xc8, ret}, make([]byte, 12)...)
	for i := 4; i < len(padded); i++ {
		padded[i] = padByte
	}
	if got := disassemble(padded, branchLen, true, unknownOracle); got != branchLen {
		t.Fatalf("padding must still be accepted, cutting index = %d, want %d", got, branchLen)
	}

	code := append([]byte{ret},
		0x48, 0x89, 0xe5, 0x48, 0x83, 0xec, 0x10, 0x48, 0x85, 0xc0, 0x75, 0xcc)
	if err := recoverOf(func() { disassemble(code, branchLen, true, unknownOracle) }); err == nil {
		t.Fatal("real code after RET must still be refused when the boundary is unknown")
	}
}

// firstRetEnd returns the offset just past the first RET, decoding whole
// instructions the way Disassemble does.
func firstRetEnd(t *testing.T, code []byte) int {
	t.Helper()
	for pos := 0; pos < len(code); {
		in, err := x86asm.Decode(code[pos:], 64)
		if err != nil {
			t.Fatalf("decode at %d: %v", pos, err)
		}
		if in.Op == x86asm.RET {
			return pos + in.Len
		}
		pos += in.Len
	}
	t.Fatal("fixture has no RET")
	return -1
}

// TestDisassembleShortFunction covers the case where the target function ends
// before `required` bytes. Overwriting past its RET is safe only when the
// remaining bytes are linker padding rather than a neighbouring function.
func TestDisassembleShortFunction(t *testing.T) {
	// A real function prologue, used as the "neighbour" that must not be clobbered:
	// MOV RBP, RSP; SUB RSP, 0x10; TEST RAX, RAX; JNE ...
	neighbour := []byte{0x48, 0x89, 0xe5, 0x48, 0x83, 0xec, 0x10, 0x48, 0x85, 0xc0, 0x75, 0xcc}

	padded := func(prefix []byte) []byte {
		out := append([]byte{}, prefix...)
		for len(out) < 32 {
			out = append(out, padByte)
		}
		return out
	}

	tests := []struct {
		name    string
		code    []byte
		wantErr bool
	}{
		{
			name: "one byte function followed by padding",
			code: padded([]byte{ret}),
		},
		{
			// LEA RCX, [RAX+RAX*2]; LEA RCX, [RCX+7]; ADD RBX, RAX; RET at offset 10.
			name: "function ending just before the branch width",
			code: padded([]byte{0x48, 0x8d, 0x0c, 0x40, 0x48, 0x8d, 0x49, 0x07, 0x48, 0x01, 0xc3}),
		},
		{
			name:    "one byte function followed by a neighbour",
			code:    append([]byte{ret}, neighbour...),
			wantErr: true,
		},
		{
			// MOV RAX, RCX; RET; two bytes of padding; then a neighbour that
			// still starts inside the window we would overwrite.
			name:    "padding that ends before the branch width",
			code:    append([]byte{0x48, 0x89, 0xc8, ret, padByte, padByte}, neighbour...),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := recoverOf(func() { Disassemble(tt.code, branchLen, true) })
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected Disassemble to refuse, but it returned")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected Disassemble to accept, got %v", err)
			}
			// The cutting index must span the whole branch sequence, otherwise
			// Unpatch would restore fewer bytes than Patch overwrote.
			if got := Disassemble(tt.code, branchLen, true); got != branchLen {
				t.Fatalf("cutting index = %d, want %d", got, branchLen)
			}
		})
	}
}

// TestDisassembleShortFunctionUnsafe pins the OptUnsafe path: when the caller
// opts out of the length check, Disassemble must not refuse for any reason.
func TestDisassembleShortFunctionUnsafe(t *testing.T) {
	code := []byte{ret, 0x48, 0x89, 0xe5, 0x48, 0x83, 0xec, 0x10, 0x48, 0x85, 0xc0, 0x75, 0xcc}
	if err := recoverOf(func() { Disassemble(code, branchLen, false) }); err != nil {
		t.Fatalf("unchecked path must not refuse, got %v", err)
	}
}

// TestDisassembleReturnsInstructionBoundary is the regression test for the
// cutting index landing in the middle of an instruction.
//
// patch.go uses the return value both as the length of the prefix copied into
// the trampoline and as the address it branches back to (targetAddr+cuttingIdx),
// so a value that is not an instruction boundary produces a trampoline whose
// last instruction is truncated and whose jump lands mid-instruction. Returning
// the raw `required` from the RET branch did exactly that whenever the bytes
// after the RET were real code rather than 0xCC padding, which is reachable via
// MockUnsafe (checkLen=false).
func TestDisassembleReturnsInstructionBoundary(t *testing.T) {
	tests := []struct {
		name string
		code []byte
		want int
	}{
		{
			// The shape that broke orgdraft: a frameless function with an early
			// return, whose continuation is a 4-byte IMUL straddling `required`.
			//   @0 TEST RAX,RAX (3) | @3 JLE +6 (2) | @5 MOV EAX,7 (5)
			//   @10 RET (1)         | @11 IMUL RAX,RAX (4) -> covers [11,15)
			// The old code returned 12, i.e. two bytes into that IMUL.
			name: "early ret followed by a wide instruction",
			code: []byte{
				0x48, 0x85, 0xc0, // TEST RAX, RAX
				0x7e, 0x06, // JLE +6
				0xb8, 0x07, 0x00, 0x00, 0x00, // MOV EAX, 7
				ret,                    // RET   @10
				0x48, 0x0f, 0xaf, 0xc0, // IMUL RAX, RAX  @11..14
				0x48, 0x83, 0xc0, 0x03, // ADD RAX, 3
				ret,
			},
			want: 15,
		},
		{
			// Padding is INT3, one byte wide, so rounding up cannot overshoot:
			// the boundary is exactly `required`, same value the old code returned.
			name: "early ret followed by padding",
			code: []byte{0x48, 0x89, 0xc8, ret, padByte, padByte, padByte, padByte,
				padByte, padByte, padByte, padByte, padByte, padByte, padByte, padByte},
			want: branchLen,
		},
		{
			// A single 2-byte instruction straddling the branch width, so the
			// boundary sits one byte past `required` rather than on it.
			name: "early ret followed by a two byte instruction",
			code: []byte{
				0x48, 0x85, 0xc0, // TEST RAX, RAX
				0x7e, 0x06, // JLE +6
				0xb8, 0x07, 0x00, 0x00, 0x00, // MOV EAX, 7
				ret,        // RET  @10
				0x31, 0xc0, // XOR EAX, EAX  @11..12
				ret,
			},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// checkLen=false is the MockUnsafe path, the only one that is allowed
			// to walk past real (non-padding) code.
			got := Disassemble(tt.code, branchLen, false)
			if got != tt.want {
				t.Fatalf("cutting index = %d, want %d", got, tt.want)
			}
			if got < branchLen {
				t.Fatalf("cutting index %d is shorter than the branch sequence %d; "+
					"Unpatch would restore fewer bytes than Patch overwrote", got, branchLen)
			}
			if !onInstructionBoundary(t, tt.code, got) {
				t.Fatalf("cutting index %d is not an instruction boundary", got)
			}
		})
	}
}

// onInstructionBoundary decodes the byte stream from the top and reports whether
// idx is reachable by whole instructions, which is the property patch.go needs.
func onInstructionBoundary(t *testing.T, code []byte, idx int) bool {
	t.Helper()
	for pos := 0; pos < len(code); {
		if pos == idx {
			return true
		}
		if pos > idx {
			return false
		}
		in, err := x86asm.Decode(code[pos:], 64)
		if err != nil {
			t.Fatalf("decode at %d: %v", pos, err)
		}
		pos += in.Len
	}
	return false
}

func recoverOf(f func()) (err interface{}) {
	defer func() { err = recover() }()
	f()
	return nil
}

// TestWithinTargetFuncOnRealCode exercises the real pclntab oracle rather than a
// stand-in, because everything above injects its own boundary and would keep
// passing even if withinTargetFunc were wrong.
//
// It asserts the three properties Disassemble leans on:
//   - a live function entry is recognised as such,
//   - the tail padding still belongs to the function (that is the whole reason
//     pclntab is usable here: the padding carries no symbol of its own),
//   - a pc that is not text at all reports "unknown" rather than lying, which is
//     what makes the padding fallback reachable.
func TestWithinTargetFuncOnRealCode(t *testing.T) {
	entry := reflect.ValueOf(sampleGetter).Pointer()
	// The window must outrun the function itself, otherwise the "past the slot"
	// probe below would still land inside it. Compiled with -N the same source is
	// several times larger than with optimisations on, so give it room.
	const window = 4096
	code := common.BytesOf(entry, window)

	if within, ok := withinTargetFunc(code, 0); !ok || !within {
		t.Fatalf("own entry: within=%v ok=%v, want true/true", within, ok)
	}

	// Walk out to the last byte the runtime still attributes to this function and
	// check that everything from the code's end to there is padding, i.e. that the
	// slot really is "my code + my padding".
	fn := runtime.FuncForPC(entry)
	last := -1
	for off := 1; off < window; off++ {
		g := runtime.FuncForPC(entry + uintptr(off))
		if g == nil || g.Entry() != fn.Entry() {
			break
		}
		last = off
	}
	if last < 0 {
		t.Fatal("could not find any offset inside the function")
	}
	if last >= window-1 {
		t.Fatalf("function does not end within the %d byte window; widen it", window)
	}
	if within, ok := withinTargetFunc(code, last); !ok || !within {
		t.Fatalf("offset %d inside the slot: within=%v ok=%v, want true/true", last, within, ok)
	}
	if within, ok := withinTargetFunc(code, last+1); ok && within {
		t.Fatalf("offset %d is past the slot but reported as inside", last+1)
	}

	// A heap slice is not text, so the lookup must decline to answer.
	heap := make([]byte, 32)
	if _, ok := withinTargetFunc(heap, 0); ok {
		t.Fatal("a non-text pointer must report ok=false, not a verdict")
	}
	if _, ok := withinTargetFunc(nil, 0); ok {
		t.Fatal("an empty slice must report ok=false")
	}
}

type sampleRecv struct{ v int32 }

// sampleGetter is shaped like the field failures: nil guard, early RET, second
// return path. go:noinline keeps it a real standalone function even without -l.
//
//go:noinline
func sampleGetter(r *sampleRecv) int32 {
	if r != nil {
		return r.v
	}
	return 0
}
