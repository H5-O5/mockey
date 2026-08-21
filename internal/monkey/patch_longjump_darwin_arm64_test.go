//go:build darwin && arm64

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
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/bytedance/mockey/internal/monkey/common"
	"github.com/bytedance/mockey/internal/monkey/inst"
)

var arm64LongJumpGlobal = 41

//go:noinline
func arm64LongJumpProbe(value *int) int {
	if value == nil {
		return arm64LongJumpGlobal
	}
	return *value
}

func usesLongJumpPath(fn interface{}) bool {
	addr := reflect.ValueOf(fn).Pointer()
	return !inst.TooShortToPatch(common.BytesOf(addr, 64), len(inst.BranchInto(0)))
}

func displacedPrefixHasPCRelative(fn interface{}) bool {
	addr := reflect.ValueOf(fn).Pointer()
	entry := common.BytesOf(addr, 64)
	cut, ok := inst.TryDisassemble(entry, len(inst.BranchInto(0)), true)
	if !ok {
		return false
	}
	for pos := 0; pos < cut; pos += 4 {
		if arm64EncodingIsPCRelative(entry[pos:]) {
			return true
		}
	}
	return false
}

func arm64EncodingIsPCRelative(code []byte) bool {
	if len(code) < 4 {
		return false
	}
	enc := *(*uint32)(unsafe.Pointer(&code[0]))
	if enc&0x1f000000 == 0x10000000 { // ADR / ADRP
		return true
	}
	if enc&0x7c000000 == 0x14000000 { // B / BL <label>
		return true
	}
	if enc&0x7e000000 == 0x34000000 { // CBZ / CBNZ
		return true
	}
	if enc&0x7e000000 == 0x36000000 { // TBZ / TBNZ
		return true
	}
	if enc&0xff000010 == 0x54000000 { // B.cond
		return true
	}
	if enc&0x3b000000 == 0x18000000 { // LDR literal
		return true
	}
	return false
}

// The darwin/arm64 regression is process-fatal on the broken code path, so the
// test probes it in a subprocess: before the fix the child dies with
// "unexpected return pc"/"unknown caller pc"; after the fix origin executes
// normally and the child exits cleanly.
func TestArm64LongJumpRelocatesPCRelativePrefix(t *testing.T) {
	if !usesLongJumpPath(arm64LongJumpProbe) {
		t.Skip("probe is still short in this build; run with -gcflags='all=-N -l' to force the long-jump path")
	}
	if !displacedPrefixHasPCRelative(arm64LongJumpProbe) {
		t.Skip("probe entry is not PC-relative in this build")
	}

	if os.Getenv("MOCKEY_ARM64_LONGJUMP_CHILD") == "1" {
		runArm64LongJumpProbe(t)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestArm64LongJumpRelocatesPCRelativePrefix$")
	cmd.Env = append(os.Environ(), "MOCKEY_ARM64_LONGJUMP_CHILD=1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("long-jump relocation probe failed in subprocess: %v\n%s", err, output)
	}
	if strings.Contains(string(output), "fatal error: unknown caller pc") ||
		strings.Contains(string(output), "unexpected return pc") {
		t.Fatalf("subprocess still hit the legacy long-jump crash:\n%s", output)
	}
}

func runArm64LongJumpProbe(t *testing.T) {
	var proxy func(*int) int
	patch := PatchFunc(arm64LongJumpProbe, func(*int) int { return -1 }, &proxy, false)
	patched := true
	defer func() {
		if patched {
			patch.Unpatch()
		}
	}()

	if got := arm64LongJumpProbe(nil); got != -1 {
		t.Fatalf("patched target = %d, want -1", got)
	}
	value := 7
	if got := proxy(&value); got != value {
		t.Fatalf("origin pointer branch = %d, want %d", got, value)
	}
	if got := proxy(nil); got != arm64LongJumpGlobal {
		t.Fatalf("origin global branch = %d, want %d", got, arm64LongJumpGlobal)
	}

	patch.Unpatch()
	patched = false
	if got := arm64LongJumpProbe(nil); got != arm64LongJumpGlobal {
		t.Fatalf("unpatched target = %d, want %d", got, arm64LongJumpGlobal)
	}
}
