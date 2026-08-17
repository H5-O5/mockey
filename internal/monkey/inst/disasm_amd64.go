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
	"reflect"
	"runtime"
	"unsafe"

	"github.com/bytedance/mockey/internal/monkey/common"
	"github.com/bytedance/mockey/internal/tool"
	"golang.org/x/arch/x86/x86asm"
)

func calcFnAddrRange(name string, fn func()) (uintptr, uintptr) {
	v := reflect.ValueOf(fn)
	var start, end uintptr
	start = v.Pointer()
	maxScan := 2000
	code := common.BytesOf(v.Pointer(), 2000)
	pos := 0

	for pos < maxScan {
		inst, err := x86asm.Decode(code[pos:], 64)
		tool.Assert(err == nil, err)

		args := []interface{}{name, inst.Op}
		for i := range inst.Args {
			args = append(args, inst.Args[i])
		}

		if inst.Op == x86asm.RET {
			end = start + uintptr(pos)
			return start, end
		}

		pos += int(inst.Len)
	}
	tool.Assert(false, "%v ret not found", name)
	return 0, 0
}

// padByte is what the amd64 linker writes between functions.
// See cmd/link/internal/amd64/obj.go: CodePad = []byte{0xCC} (INT $3),
// emitted to pad each function entry up to funcAlign = 32
// (cmd/link/internal/amd64/l.go).
const padByte = 0xcc

// withinTargetFunc reports whether the byte at code[offset] still belongs to the
// function that code starts at, i.e. whether overwriting it can only ever damage
// the target itself and never a neighbour.
//
// It answers with pclntab, via the public runtime.FuncForPC. The lookup is a
// binary search over the function address table, so *every* pc in
// [entry_i, entry_{i+1}) maps to function i -- including the alignment padding
// at the tail, which has no symbol of its own. That is precisely the question we
// need answered: "is this address still inside my function's slot?", where the
// slot is code plus its padding, and it is exactly the region the next function
// does not occupy.
//
// ok=false means "cannot tell", not "no": code may not point into the text
// segment at all (unit tests pass ordinary heap slices), and FuncForPC returns
// nil for a pc outside every known function. Callers must treat !ok as "no new
// information" and fall back to the padding rule, so that a failed lookup can
// never turn a patch that used to work into a refusal.
func withinTargetFunc(code []byte, offset int) (within bool, ok bool) {
	if len(code) == 0 {
		return false, false
	}
	// code is common.BytesOf(targetAddr, bufSize), so its first byte *is* the
	// function entry the caller is about to patch.
	entry := uintptr(unsafe.Pointer(&code[0]))
	self := runtime.FuncForPC(entry)
	if self == nil || self.Entry() != entry {
		// Not a known function entry: either not text at all, or entry sits in
		// the middle of one (hand-written assembly sublabels do that). Either way
		// we have no trustworthy slot to compare against.
		return false, false
	}
	at := runtime.FuncForPC(entry + uintptr(offset))
	if at == nil {
		// Past the last function of the text segment, so there is no neighbour
		// this far out either -- but say "unknown" rather than guess.
		return false, false
	}
	return at.Entry() == self.Entry(), true
}

func Disassemble(code []byte, required int, checkLen bool) int {
	return disassemble(code, required, checkLen, withinTargetFunc)
}

// disassemble is Disassemble with the function-boundary oracle injected. The real
// oracle needs code to point at a live function entry, which a test cannot
// fabricate for a synthetic byte slice, so tests substitute their own to drive
// both answers deterministically.
func disassemble(code []byte, required int, checkLen bool, within func([]byte, int) (bool, bool)) int {
	var pos int
	var err error
	var inst x86asm.Inst

	for pos < required {
		inst, err = x86asm.Decode(code[pos:], 64)
		tool.Assert(err == nil, err)
		tool.DebugPrintf("Disassemble: inst: %v\n", inst)
		if inst.Op == x86asm.RET {
			// A return, not necessarily *the* return: the target may well continue
			// past it. Either way the branch sequence we must write is longer than
			// what we have walked so far, so we have to keep going and decide what
			// the bytes after this RET are.
			//
			// Overwriting them is fatal only if they belong to the *next* function.
			// Everything up to the next function's entry is fair game, and that
			// covers two different kinds of bytes:
			//
			//  1. Linker padding. The amd64 linker aligns each function entry and
			//     fills the gap with 0xCC, dead space no control flow ever enters.
			//
			//  2. More of the target's own code, i.e. a second return path. This is
			//     the common case for the tiny nil-guarded getters Go generates:
			//         (*BaseResp).GetStatusCode, 11 code bytes
			//           +0  48 85 c0  TEST RAX, RAX
			//           +3  74 03     JE +3
			//           +5  8b 00     MOV EAX, [RAX]
			//           +7  c3        RET          <- first RET, ends at +8
			//           +8  31 c0     XOR EAX, EAX <- nil branch, real code
			//           +10 c3        RET
			//           +11 cc ...    padding starts here
			//     Clobbering +8..+10 is safe *because we are patching the entry*:
			//     the first thing we write is an unconditional branch to the hook,
			//     so the target's own instructions are unreachable, all of them.
			//     Nothing can jump into the middle of the function either -- Go
			//     emits no such edges across a function boundary, and the only
			//     inbound edge that survives the patch is the trampoline's, which
			//     lands on the boundary we return here. So a second return path is
			//     no more live than the padding next to it.
			//
			// Hence the rule is not "stop at the first RET" but "stay inside the
			// target's own slot". We accept a byte when either
			//   (a) pclntab says the address is still in the target function
			//       (authoritative, and it covers the padding too), or
			//   (b) it decodes as the 0xCC pad, which no function starts with, so
			//       it cannot be a neighbour's first instruction.
			// (b) is kept as the fallback for when the lookup cannot answer, so a
			// nil FuncForPC never downgrades a patch that used to succeed. Note
			// 0xCC has to be tested at an *instruction boundary*: scanning raw
			// bytes would misread the 0xCC in e.g. "75 cc" (JNE -52) as padding.
			//
			// Why we keep decoding instead of `return required` -- read this before
			// "simplifying" it back:
			//
			// The value returned here is PatchValue's cuttingIdx, and it is used for
			// three things at once (internal/monkey/patch.go):
			//   1. the trampoline saves code[:cuttingIdx],
			//   2. the trampoline ends with BranchTo(targetAddr+cuttingIdx),
			//   3. Unpatch restores cuttingIdx bytes.
			// (2) makes cuttingIdx a *jump target in the original instruction
			// stream*, so it MUST sit on an instruction boundary; (1) means the
			// saved prefix must end on one too, or the trampoline's last "instruction"
			// is a truncated fragment. `required` is just the length of the branch
			// sequence (12 bytes here) -- nothing makes it land on a boundary.
			//
			// The non-RET exit below returns `pos`, which is boundary-aligned by
			// construction. The RET branch returning the raw `required` was the odd
			// one out, and it is exactly the bug: for a frameless function with an
			// early `return`, e.g.
			//     @0 TEST(3) @3 JLE(2) @5 MOV(5) @10 RET(1) @11 IMUL(4) @15 ...
			// the loop hits RET at pos=10 and returns 12 -- which is the *middle* of
			// the 4-byte IMUL at [11,15). The trampoline then keeps only the orphan
			// REX prefix 0x48 of that IMUL and jumps back into its middle, so the CPU
			// executes a garbage instruction stream. Depending on how those bytes
			// happen to decode you get a SIGSEGV or, worse, a silently wrong result.
			//
			// So instead of trusting `required`, walk whole instructions until we are
			// at or past it, and return that boundary. The returned value is still
			// >= required (never restoring/overwriting less than the branch sequence),
			// it is just rounded up to the next boundary instead of cutting blindly.
			pos += inst.Len // step over the RET itself first -- otherwise the very
			// first iteration would test the RET opcode 0xc3 against padByte and
			// wrongly report "function is too short to patch", and would re-decode
			// the same RET forever.
			for pos < required {
				inst, err = x86asm.Decode(code[pos:], 64)
				tool.Assert(err == nil, err)
				if checkLen {
					// Ask about the last byte this instruction occupies, not its
					// first: an instruction that starts inside the target must also
					// end inside it, or writing the branch would still reach into
					// the neighbour.
					end := pos + inst.Len
					inside, known := within(code, end-1)
					tool.Assert(inside || (!known && code[pos] == padByte),
						// Keep the historical wording so existing reports and
						// searches still match, but say what actually went wrong:
						// the branch sequence would not fit without writing over
						// whatever follows the target.
						"function is too short to patch: writing %v bytes would run past the "+
							"target function into the next one at offset %v", required, pos)
				}
				pos += inst.Len
			}
			// Bound check for the caller: pos < required <= len(hookCode) == 12 on
			// entry to the last iteration, and an x86 instruction is at most 15 bytes,
			// so pos <= 26 here -- well inside the 64-byte targetCodeBuf and nowhere
			// near the page the trampoline is written into.
			return pos
		}
		pos += inst.Len
	}
	return pos
}

func GetGenericJumpAddr(addr uintptr, maxScan uint64) uintptr {
	code := common.BytesOf(addr, int(maxScan))
	var pos uint64
	var err error
	var inst x86asm.Inst

	allAddrs := []uintptr{}
	for pos < maxScan {
		inst, err = x86asm.Decode(code[pos:], 64)
		tool.Assert(err == nil, err)

		args := []interface{}{inst.Op}
		for i := range inst.Args {
			args = append(args, inst.Args[i])
		}
		tool.DebugPrintf("%v\t%v\t%v\t%v\t%v\t%v\n", args...)

		if inst.Op == x86asm.RET {
			break
		}

		if inst.Op == x86asm.CALL {
			rel := int32(inst.Args[0].(x86asm.Rel))
			fnAddr := calcAddr(uintptr(unsafe.Pointer(&code[0]))+uintptr(pos+uint64(inst.Len)), rel)
			isExtraCall, extraName := isGenericProxyCallExtra(fnAddr)
			tool.DebugPrintf("found CALL, raw is: %x, rel: %v,  raw is: %x,  fnAddr: %v, isExtraCall: %v, extraName: %v\n", inst.String(), rel, fnAddr, isExtraCall, extraName)
			if !isExtraCall {
				allAddrs = append(allAddrs, fnAddr)
			}
		}
		pos += uint64(inst.Len)
	}
	tool.Assert(len(allAddrs) == 1, "invalid callAddr: %v", allAddrs)
	return allAddrs[0]
}

func calcAddr(from uintptr, rel int32) uintptr {
	tool.DebugPrintf("calc CALL addr, from: %x(%v) CALL: %x\n", from, from, rel)

	var dest uintptr
	if rel < 0 {
		dest = from - uintptr(uint32(-rel))
	} else {
		dest = from + uintptr(rel)
	}

	tool.DebugPrintf("L->H:%v rel: %v from: %x(%v) dest: %x(%v), distance: %v\n", rel > 0, rel, from, from, dest, dest, from-dest)
	return dest
}
