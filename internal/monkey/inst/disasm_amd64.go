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

func Disassemble(code []byte, required int, checkLen bool) int {
	var pos int
	var err error
	var inst x86asm.Inst

	for pos < required {
		inst, err = x86asm.Decode(code[pos:], 64)
		tool.Assert(err == nil, err)
		tool.DebugPrintf("Disassemble: inst: %v\n", inst)
		if inst.Op == x86asm.RET {
			// The target's own code ends here, so it is shorter than the branch
			// sequence we need to write. That is only fatal if the remaining bytes
			// belong to the *next* function. On amd64 they do not: the linker aligns
			// every function entry to 32 bytes and fills the gap with 0xCC, which is
			// dead space no control flow ever enters. Overwriting it is harmless, so
			// only refuse when the tail is not padding.
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
			//
			// Note this changes nothing for the checkLen==true (Mock) path: there
			// every skipped byte must be 0xCC padding, and INT3 is 1 byte long, so pos
			// lands exactly on `required` anyway. Only the checkLen==false
			// (MockUnsafe) path, which is allowed to run past real code, is fixed.
			pos += inst.Len // step over the RET itself first -- otherwise the very
			// first iteration would test the RET opcode 0xc3 against padByte and
			// wrongly report "function is too short to patch", and would re-decode
			// the same RET forever.
			for pos < required {
				tool.Assert(code[pos] == padByte || !checkLen, "function is too short to patch")
				inst, err = x86asm.Decode(code[pos:], 64)
				tool.Assert(err == nil, err)
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
