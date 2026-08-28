//go:build windows

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

package fn

import (
	"syscall"
	"unsafe"
)

var (
	modkernel32             = syscall.NewLazyDLL("kernel32.dll")
	procVirtualQuery        = modkernel32.NewProc("VirtualQuery")
	memCommit        uint32 = 0x1000
	pageReadAccess   uint32 = 0x02 | 0x04 | 0x20 | 0x40 // READONLY, READWRITE, EXECUTE_READ, EXECUTE_READWRITE
)

type memoryBasicInformation struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	PartitionId       uint16
	_                 uint16
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

// derefUintptrSafe reads *uintptr at p only if the address is committed and
// readable. See the unix implementation for why this guard exists.
func derefUintptrSafe(p uintptr) (uintptr, bool) {
	var mbi memoryBasicInformation
	ret, _, _ := procVirtualQuery.Call(
		uintptr(unsafe.Pointer(p)),
		uintptr(unsafe.Pointer(&mbi)),
		unsafe.Sizeof(mbi),
	)
	if ret == 0 {
		return 0, false
	}
	if mbi.State&memCommit == 0 {
		return 0, false
	}
	if mbi.Protect&pageReadAccess == 0 {
		return 0, false
	}
	return *(*uintptr)(unsafe.Pointer(p)), true
}
