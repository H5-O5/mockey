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

const errFunctionTooShort = "function is too short to patch"

func TryDisassemble(code []byte, required int, checkLen bool) (pos int, ok bool) {
	defer func() {
		if err := recover(); err != nil {
			if err == errFunctionTooShort {
				pos, ok = 0, false
				return
			}
			panic(err)
		}
	}()
	return Disassemble(code, required, checkLen), true
}
