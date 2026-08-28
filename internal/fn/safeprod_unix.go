//go:build unix

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
	"unsafe"

	"golang.org/x/sys/unix"
)

// derefUintptrSafe reads *uintptr at p only if the underlying page(s) are
// mapped and resident. It returns the value and true on success, or zero and
// false when the address cannot be safely read.
//
// This exists because GenericInfo values come from instruction analysis
// (inst.GetGenericAddr) and from the live register/argument state of a patched
// generic call. Either source can, for wrappers the analyzer does not fully
// model, hand back a non-nil address that is not a dictionary. Dereferencing
// such an address raises SIGSEGV, which Go cannot recover. A SIGSEGV inside a
// test mock aborts the whole test process, which is strictly worse than
// treating the two dictionaries as not equal (the mock falls through to the
// original function).
func derefUintptrSafe(p uintptr) (uintptr, bool) {
	pageSize := uintptr(unix.Getpagesize())
	base := p &^ (pageSize - 1)
	// Addresses in the null page are never valid dictionaries.
	if base == 0 {
		return 0, false
	}
	// GenericInfo values are word-aligned, so the 8-byte read stays within the
	// single page containing p. mincore reports that page's residency without
	// touching the memory itself.
	vec := make([]byte, 1)
	_, _, errno := unix.RawSyscall(
		unix.SYS_MINCORE,
		base,
		pageSize,
		uintptr(unsafe.Pointer(&vec[0])),
	)
	if errno != 0 || vec[0]&1 == 0 {
		return 0, false
	}
	return *(*uintptr)(unsafe.Pointer(p)), true
}
