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
	"encoding/binary"
	"unsafe"
)

const shortBranchSize = 5

func BranchTo(to uintptr) (res []byte) {
	res = append(res, rdxMOV(to)...)         // MOVABS RDX, to
	res = append(res, []byte{0xff, 0xe2}...) // JMP RDX
	return
}

func BranchInto(to uintptr) (res []byte) {
	res = append(res, rdxMOV(to)...)         // MOVABS RDX, to
	res = append(res, []byte{0xff, 0x22}...) // JMP [RDX]
	return
}

func ShortBranchSize() int {
	return shortBranchSize
}

func BranchIntoShort(from, to uintptr) ([]byte, bool) {
	delta, ok := subtractAddress(to, from+shortBranchSize)
	if !ok || delta < -1<<31 || delta > 1<<31-1 {
		return nil, false
	}
	result := make([]byte, shortBranchSize)
	result[0] = 0xe9
	binary.LittleEndian.PutUint32(result[1:], uint32(int32(delta)))
	return result, true
}

// rdxMOV moves the 64bit value to rdx register, using the following instruction:
// MOVABS RDX, val
func rdxMOV(val uintptr) []byte {
	res := make([]byte, unsafe.Sizeof(val))
	*(*uintptr)(unsafe.Pointer(&res[0])) = val
	res = append([]byte{0x48, 0xba}, res...)
	return res
}
