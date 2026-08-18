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

// TryDisassemble reports whether Disassemble can find a cutting point, without
// panicking when it cannot.
//
// Disassemble rejects by panicking, and it has more than one reason to do so:
// "function is too short to patch" when the tail is not 0xCC padding, and
// "unrecognized instruction" (from x86asm.Decode) when the window starts or
// ends in the middle of something it cannot decode. Every one of those means
// the same thing to a caller that is only probing -- "this cutting point is not
// usable" -- so all of them must turn into ok=false rather than escaping.
//
// Recovering only the too-short case is a real bug: the short-jump path probes
// a 5-byte window, which lands mid-instruction far more often than the 12-byte
// one, so "unrecognized instruction" escaped and crashed tests that the legacy
// path had rejected cleanly. Probing must never be able to fail louder than the
// path it is probing for.
func TryDisassemble(code []byte, required int, checkLen bool) (pos int, ok bool) {
	defer func() {
		if err := recover(); err != nil {
			pos, ok = 0, false
		}
	}()
	return Disassemble(code, required, checkLen), true
}
