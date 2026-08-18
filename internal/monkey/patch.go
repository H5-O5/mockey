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

package monkey

import (
	"reflect"

	"github.com/bytedance/mockey/internal/monkey/common"
	"github.com/bytedance/mockey/internal/monkey/fn"
	"github.com/bytedance/mockey/internal/monkey/inst"
	"github.com/bytedance/mockey/internal/monkey/mem"
	"github.com/bytedance/mockey/internal/tool"
)

// Patch is a context that holds the address and original codes of the patched function.
type Patch struct {
	code     []byte
	original []byte
	base     uintptr
}

// Unpatch restores the patched function to the original function.
func (p *Patch) Unpatch() {
	mem.WriteWithSTW(p.base, p.original)
	common.ReleasePage(p.code)
}

// PatchValue replace the target function with a hook function, and stores the target function in the proxy function
// for future restore. Target and hook are values of function. Proxy is a value of proxy function pointer.
func PatchValue(target, hook, proxy reflect.Value, unsafe, generic bool) *Patch {
	tool.Assert(hook.Kind() == reflect.Func, "'%s' is not a function", hook.Kind())
	tool.Assert(proxy.Kind() == reflect.Ptr, "'%v' is not a function pointer", proxy.Kind())

	targetAddr := target.Pointer()
	if generic {
		// we assume that generic call/bl op is located in first 200 bytes of codes from targetAddr
		targetAddr = inst.GetGenericJumpAddr(targetAddr, 10000)
	}
	// The first few bytes of the target function code
	const bufSize = 64
	targetCodeBuf := common.BytesOf(targetAddr, bufSize)
	hookAddr := common.PtrAt(hook)
	hookCode := inst.BranchInto(hookAddr)
	var cuttingIdx int
	var proxyCode []byte
	var proxyPrefix []byte

	if unsafe || inst.ShortBranchSize() == 0 {
		// Keep MockUnsafe and non-amd64 architectures on the legacy path.
		cuttingIdx = inst.Disassemble(targetCodeBuf, len(hookCode), !unsafe)
		proxyCode = common.AllocatePage()
		proxyPrefix = targetCodeBuf[:cuttingIdx]
	} else if idx, ok := inst.TryDisassemble(targetCodeBuf, len(hookCode), true); ok && !inst.HasPCRelative(targetCodeBuf[:idx]) {
		// Preserve the existing 12-byte entry sequence when its trampoline prefix
		// is position independent.
		cuttingIdx = idx
		proxyCode = common.AllocatePage()
		proxyPrefix = targetCodeBuf[:cuttingIdx]
	} else {
		// A five-byte E9 reaches a nearby relay. The relay restores the function
		// value pointer in RDX before entering the reflect-generated hook.
		var ok bool
		cuttingIdx, ok = inst.TryDisassemble(targetCodeBuf, inst.ShortBranchSize(), true)
		tool.Assert(ok, "function is too short to patch")

		var err error
		proxyCode, err = common.AllocatePageNear(targetAddr)
		tool.Assert(err == nil, "allocate near trampoline failed: %v", err)
		proxyAddr := common.PtrOf(proxyCode)
		proxyPrefix, err = inst.Relocate(targetCodeBuf[:cuttingIdx], targetAddr, proxyAddr)
		if err != nil {
			common.ReleasePage(proxyCode)
			tool.Assert(false, "relocate trampoline failed: %v", err)
		}

		branchBack := inst.BranchTo(targetAddr + uintptr(cuttingIdx))
		relayOffset := align(len(proxyPrefix)+len(branchBack), 16)
		relayCode := inst.BranchInto(hookAddr)
		tool.Assert(relayOffset+len(relayCode) <= len(proxyCode), "relay does not fit in proxy page")
		copy(proxyCode[relayOffset:], relayCode)
		hookCode, ok = inst.BranchIntoShort(targetAddr, proxyAddr+uintptr(relayOffset))
		if !ok {
			common.ReleasePage(proxyCode)
			tool.Assert(false, "near trampoline is outside rel32 range")
		}
	}

	original := append([]byte(nil), targetCodeBuf[:cuttingIdx]...)
	// cuttingIdx is no longer guaranteed to equal len(hookCode): when the target
	// ends in an early RET, Disassemble rounds up to the next instruction boundary
	// past it, so cuttingIdx may exceed len(hookCode) (bounded by 12+15-1 = 26 on
	// amd64, one branch width plus at most one maximal x86 instruction). Both the
	// slice read below and the trampoline write must still fit, so pin the two
	// invariants rather than assume them: targetCodeBuf is bufSize (64) bytes, and
	// a page is at least 4096 bytes, so neither can fire today -- this is here to
	// fail loudly instead of corrupting memory if either bound ever changes.
	tool.Assert(cuttingIdx <= bufSize, "cutting index %v exceeds the %v bytes read from the target", cuttingIdx, bufSize)
	tool.Assert(len(proxyPrefix)+len(inst.BranchTo(0)) <= len(proxyCode), "trampoline (%v code bytes + branch) does not fit in a %v byte page", len(proxyPrefix), len(proxyCode))
	// save the original code before the cutting point
	copy(proxyCode, proxyPrefix)
	// construct the branch instruction, i.e. jump to the cutting point
	copy(proxyCode[len(proxyPrefix):], inst.BranchTo(targetAddr+uintptr(cuttingIdx)))
	// inject the proxy code to the proxy function
	fn.InjectInto(proxy, proxyCode)

	tool.DebugPrintf("PatchValue: hook code len(%v), cuttingIdx(%v)\n", len(hookCode), cuttingIdx)

	// replace target function codes before the cutting point
	mem.WriteWithSTW(targetAddr, hookCode)

	return &Patch{base: targetAddr, code: proxyCode, original: original}
}

func align(value, alignment int) int {
	return (value + alignment - 1) &^ (alignment - 1)
}

func PatchFunc(fn, hook, proxy interface{}, unsafe bool) *Patch {
	vv := reflect.ValueOf(fn)
	tool.Assert(vv.Kind() == reflect.Func, "'%v' is not a function", fn)
	return PatchValue(vv, reflect.ValueOf(hook), reflect.ValueOf(proxy), unsafe, false)
}
