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
	"testing"
)

// branchLen is the width of the branch sequence Disassemble is asked to make
// room for on amd64: MOV RDX, imm64 (10 bytes) + JMP RDX (2 bytes).
// It matches len(BranchInto(...)), which is what patch.go passes as `required`.
const branchLen = 12

// ret is a single-byte RET (0xc3), i.e. the shortest possible function body.
const ret = 0xc3

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

func recoverOf(f func()) (err interface{}) {
	defer func() { err = recover() }()
	f()
	return nil
}
