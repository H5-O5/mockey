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
	"fmt"
	"unsafe"

	"golang.org/x/arch/arm64/arm64asm"
)

func HasPCRelative(code []byte) bool {
	if len(code)%instLen != 0 {
		return true
	}
	for pos := 0; pos < len(code); pos += instLen {
		inst, enc, err := decodeArm64Inst(code[pos:])
		if err != nil {
			return true
		}
		if kind, _ := classifyArm64PCRelative(enc, inst); kind != arm64RelocNone {
			return true
		}
	}
	return false
}

func Relocate(code []byte, oldAddr, newAddr uintptr) ([]byte, error) {
	decoded, newSize, err := decodeArm64ForRelocation(code)
	if err != nil {
		return nil, err
	}
	offsets := make(map[int]int, len(decoded))
	for _, item := range decoded {
		offsets[item.oldOffset] = item.newOffset
	}

	result := make([]byte, newSize)
	for _, item := range decoded {
		enc, err := relocateArm64Inst(item, oldAddr, newAddr, len(code), offsets)
		if err != nil {
			return nil, fmt.Errorf("relocate %s at offset %d: %w", item.name, item.oldOffset, err)
		}
		*(*uint32)(unsafe.Pointer(&result[item.newOffset])) = enc[0]
		for i := 1; i < len(enc); i++ {
			*(*uint32)(unsafe.Pointer(&result[item.newOffset+i*instLen])) = enc[i]
		}
	}
	return result, nil
}

type arm64RelocKind int

const (
	arm64RelocNone arm64RelocKind = iota
	arm64RelocADR
	arm64RelocADRP
	arm64RelocB
	arm64RelocBL
	arm64RelocCB
	arm64RelocTB
	arm64RelocBCond
	arm64RelocLDRLiteral
)

type decodedArm64Inst struct {
	oldOffset int
	newOffset int
	newLen    int
	inst      arm64asm.Inst
	enc       uint32
	kind      arm64RelocKind
	name      string
}

func decodeArm64ForRelocation(code []byte) ([]decodedArm64Inst, int, error) {
	if len(code)%instLen != 0 {
		return nil, 0, fmt.Errorf("arm64 relocation requires whole instructions, got %d bytes", len(code))
	}
	decoded := make([]decodedArm64Inst, 0, len(code)/instLen)
	oldOffset, newOffset := 0, 0
	for oldOffset < len(code) {
		inst, enc, err := decodeArm64Inst(code[oldOffset:])
		if err != nil {
			return nil, 0, fmt.Errorf("decode instruction at offset %d: %w", oldOffset, err)
		}
		kind, name := classifyArm64PCRelative(enc, inst)
		newLen, err := arm64RelocatedSize(kind, name)
		if err != nil {
			return nil, 0, err
		}
		decoded = append(decoded, decodedArm64Inst{
			oldOffset: oldOffset,
			newOffset: newOffset,
			newLen:    newLen,
			inst:      inst,
			enc:       enc,
			kind:      kind,
			name:      name,
		})
		oldOffset += instLen
		newOffset += newLen
	}
	return decoded, newOffset, nil
}

func arm64RelocatedSize(kind arm64RelocKind, name string) (int, error) {
	switch kind {
	case arm64RelocNone:
		return instLen, nil
	case arm64RelocADR, arm64RelocADRP:
		return 4 * instLen, nil
	case arm64RelocB, arm64RelocBL:
		return 5 * instLen, nil
	case arm64RelocCB, arm64RelocTB, arm64RelocBCond:
		return 6 * instLen, nil
	case arm64RelocLDRLiteral:
		return 0, fmt.Errorf("%s is not yet supported on the arm64 long-jump path", name)
	default:
		return 0, fmt.Errorf("unsupported arm64 relocation kind %d", kind)
	}
}

func decodeArm64Inst(code []byte) (arm64asm.Inst, uint32, error) {
	if len(code) < instLen {
		return arm64asm.Inst{}, 0, fmt.Errorf("truncated instruction")
	}
	inst, err := arm64asm.Decode(code[:instLen])
	if err != nil {
		return arm64asm.Inst{}, 0, err
	}
	enc := *(*uint32)(unsafe.Pointer(&code[0]))
	return inst, enc, nil
}

func classifyArm64PCRelative(enc uint32, inst arm64asm.Inst) (arm64RelocKind, string) {
	switch {
	case enc&0x1f000000 == 0x10000000:
		if enc&0x80000000 != 0 {
			return arm64RelocADRP, "ADRP"
		}
		return arm64RelocADR, "ADR"
	case enc&0x7c000000 == 0x14000000:
		if enc&0x80000000 != 0 {
			return arm64RelocBL, "BL"
		}
		return arm64RelocB, "B"
	case enc&0x7e000000 == 0x34000000:
		if enc&0x01000000 != 0 {
			return arm64RelocCB, "CBNZ"
		}
		return arm64RelocCB, "CBZ"
	case enc&0x7e000000 == 0x36000000:
		if enc&0x01000000 != 0 {
			return arm64RelocTB, "TBNZ"
		}
		return arm64RelocTB, "TBZ"
	case enc&0xff000010 == 0x54000000:
		return arm64RelocBCond, "B.cond"
	case enc&0x3b000000 == 0x18000000:
		if inst.Op != 0 {
			return arm64RelocLDRLiteral, inst.Op.String() + "-literal"
		}
		return arm64RelocLDRLiteral, "LDR-literal"
	default:
		return arm64RelocNone, inst.Op.String()
	}
}

func relocateArm64Inst(item decodedArm64Inst, oldAddr, newAddr uintptr, oldBlockLen int, offsets map[int]int) ([]uint32, error) {
	if item.kind == arm64RelocNone {
		return []uint32{item.enc}, nil
	}

	oldInstAddr := oldAddr + uintptr(item.oldOffset)
	target, err := arm64Target(item.kind, item.enc, oldInstAddr)
	if err != nil {
		return nil, err
	}
	mapped, err := mapArm64Target(target, oldAddr, newAddr, oldBlockLen, offsets)
	if err != nil {
		return nil, err
	}
	if mapped != 0 {
		target = mapped
	}

	switch item.kind {
	case arm64RelocADR:
		return emitMoveAddress(target, arm64RegNumber(item.inst.Args[0])), nil
	case arm64RelocADRP:
		return emitMoveAddress(target&^0xfff, arm64RegNumber(item.inst.Args[0])), nil
	case arm64RelocB:
		return emitAbsoluteBranch(target, false), nil
	case arm64RelocBL:
		return emitAbsoluteBranch(target, true), nil
	case arm64RelocCB:
		return emitCompareAndBranch(item.enc, target), nil
	case arm64RelocTB:
		return emitTestAndBranch(item.enc, target), nil
	case arm64RelocBCond:
		return emitConditionalBranch(item.enc, target), nil
	default:
		return nil, fmt.Errorf("unsupported arm64 relocation kind %d", item.kind)
	}
}

func arm64Target(kind arm64RelocKind, enc uint32, instAddr uintptr) (uintptr, error) {
	switch kind {
	case arm64RelocADR:
		return addSignedArm64(instAddr, decodeADRImmediate(enc))
	case arm64RelocADRP:
		return addSignedArm64(instAddr&^0xfff, decodeADRImmediate(enc)<<12)
	case arm64RelocB, arm64RelocBL:
		return addSignedArm64(instAddr, decodeSignedImmediate(enc&0x03ffffff, 26)<<2)
	case arm64RelocCB, arm64RelocBCond, arm64RelocLDRLiteral:
		return addSignedArm64(instAddr, decodeSignedImmediate((enc>>5)&0x7ffff, 19)<<2)
	case arm64RelocTB:
		return addSignedArm64(instAddr, decodeSignedImmediate((enc>>5)&0x3fff, 14)<<2)
	default:
		return 0, fmt.Errorf("unsupported arm64 relocation kind %d", kind)
	}
}

func mapArm64Target(target, oldAddr, newAddr uintptr, oldBlockLen int, offsets map[int]int) (uintptr, error) {
	if target < oldAddr || target >= oldAddr+uintptr(oldBlockLen) {
		return 0, nil
	}
	mapped, ok := offsets[int(target-oldAddr)]
	if !ok {
		return 0, fmt.Errorf("relative target points into the middle of a copied instruction")
	}
	return newAddr + uintptr(mapped), nil
}

func arm64RegNumber(arg interface{}) uint32 {
	reg, ok := arg.(arm64asm.Reg)
	if !ok {
		panic(fmt.Sprintf("arm64 relocation expected register operand, got %T", arg))
	}
	switch {
	case reg >= arm64asm.W0 && reg <= arm64asm.WZR:
		return uint32(reg - arm64asm.W0)
	case reg >= arm64asm.X0 && reg <= arm64asm.XZR:
		return uint32(reg - arm64asm.X0)
	default:
		panic(fmt.Sprintf("arm64 relocation expected integer register, got %v", reg))
	}
}

func emitMoveAddress(target uintptr, reg uint32) []uint32 {
	return []uint32{
		encodeMOVZ(reg, uint16(target), 0),
		encodeMOVK(reg, uint16(target>>16), 1),
		encodeMOVK(reg, uint16(target>>32), 2),
		encodeMOVK(reg, uint16(target>>48), 3),
	}
}

func emitAbsoluteBranch(target uintptr, link bool) []uint32 {
	out := emitMoveAddress(target, uint32(x26))
	if link {
		return append(out, encodeBLR(uint32(x26)))
	}
	return append(out, encodeBR(uint32(x26)))
}

func emitCompareAndBranch(enc uint32, target uintptr) []uint32 {
	skip := uint32(6 * instLen)
	branch := encodeCBZLike(enc^0x01000000, skip)
	return append([]uint32{branch}, emitAbsoluteBranch(target, false)...)
}

func emitTestAndBranch(enc uint32, target uintptr) []uint32 {
	skip := uint32(6 * instLen)
	branch := encodeTBZLike(enc^0x01000000, skip)
	return append([]uint32{branch}, emitAbsoluteBranch(target, false)...)
}

func emitConditionalBranch(enc uint32, target uintptr) []uint32 {
	skip := uint32(6 * instLen)
	branch := encodeBCond(enc, skip, true)
	return append([]uint32{branch}, emitAbsoluteBranch(target, false)...)
}

func encodeMOVZ(reg uint32, imm uint16, shift uint32) uint32 {
	return 0xd2800000 | ((shift & 0x3) << 21) | (uint32(imm) << 5) | reg
}

func encodeMOVK(reg uint32, imm uint16, shift uint32) uint32 {
	return 0xf2800000 | ((shift & 0x3) << 21) | (uint32(imm) << 5) | reg
}

func encodeBR(reg uint32) uint32 {
	return 0xd61f0000 | ((reg & 0x1f) << 5)
}

func encodeBLR(reg uint32) uint32 {
	return 0xd63f0000 | ((reg & 0x1f) << 5)
}

func encodeCBZLike(enc uint32, offset uint32) uint32 {
	imm19 := (offset >> 2) & 0x7ffff
	enc &^= 0x7ffff << 5
	enc |= imm19 << 5
	return enc
}

func encodeTBZLike(enc uint32, offset uint32) uint32 {
	imm14 := (offset >> 2) & 0x3fff
	enc &^= 0x3fff << 5
	enc |= imm14 << 5
	return enc
}

func encodeBCond(enc uint32, offset uint32, invert bool) uint32 {
	if invert {
		enc = (enc &^ 0xf) | ((enc ^ 0x1) & 0xf)
	}
	imm19 := (offset >> 2) & 0x7ffff
	enc &^= 0x7ffff << 5
	enc |= imm19 << 5
	return enc
}

func decodeADRImmediate(enc uint32) int64 {
	immLo := (enc >> 29) & 0x3
	immHi := (enc >> 5) & 0x7ffff
	return decodeSignedImmediate((immHi<<2)|immLo, 21)
}

func decodeSignedImmediate(value uint32, bits uint) int64 {
	signBit := uint32(1) << (bits - 1)
	mask := uint32(1<<bits) - 1
	value &= mask
	if value&signBit == 0 {
		return int64(value)
	}
	return int64(value) - int64(uint32(1)<<bits)
}

func addSignedArm64(base uintptr, delta int64) (uintptr, error) {
	if delta < 0 {
		magnitude := uintptr(-delta)
		if magnitude > base {
			return 0, fmt.Errorf("pc-relative target underflows address space")
		}
		return base - magnitude, nil
	}
	result := base + uintptr(delta)
	if result < base {
		return 0, fmt.Errorf("pc-relative target overflows address space")
	}
	return result, nil
}

func rewriteADRImmediate(enc uint32, target, instAddr uintptr) (uint32, error) {
	delta, ok := subtractAddressArm64(target, instAddr)
	if !ok {
		return 0, fmt.Errorf("ADR target overflows address space")
	}
	if !fitsSignedBits(delta, 21) {
		return 0, fmt.Errorf("ADR target is outside the supported +/-1MB range")
	}
	u := uint32(delta) & 0x1fffff
	enc = enc&^(0x3<<29) | (u&0x3)<<29
	enc = enc&^(0x7ffff<<5) | ((u>>2)&0x7ffff)<<5
	return enc, nil
}

func rewriteADRPImmediate(enc uint32, target, instAddr uintptr) (uint32, error) {
	targetPage := target &^ 0xfff
	instPage := instAddr &^ 0xfff
	delta, ok := subtractAddressArm64(targetPage, instPage)
	if !ok {
		return 0, fmt.Errorf("ADRP target overflows address space")
	}
	if delta%0x1000 != 0 {
		return 0, fmt.Errorf("ADRP target is not page aligned")
	}
	scaled := delta >> 12
	if !fitsSignedBits(scaled, 21) {
		return 0, fmt.Errorf("ADRP target is outside the supported page range")
	}
	u := uint32(scaled) & 0x1fffff
	enc = enc&^(0x3<<29) | (u&0x3)<<29
	enc = enc&^(0x7ffff<<5) | ((u>>2)&0x7ffff)<<5
	return enc, nil
}

func rewriteScaledImmediate(enc uint32, target, instAddr uintptr, bits uint, mask uint32, shift uint) (uint32, error) {
	delta, ok := subtractAddressArm64(target, instAddr)
	if !ok {
		return 0, fmt.Errorf("target overflows address space")
	}
	if delta%4 != 0 {
		return 0, fmt.Errorf("target is not instruction aligned")
	}
	scaled := delta >> 2
	if !fitsSignedBits(scaled, bits) {
		return 0, fmt.Errorf("target is outside the supported range")
	}
	u := uint32(scaled) & mask
	enc = enc &^ (mask << shift)
	enc |= u << shift
	return enc, nil
}

func fitsSignedBits(value int64, bits uint) bool {
	min := -(int64(1) << (bits - 1))
	max := (int64(1) << (bits - 1)) - 1
	return value >= min && value <= max
}

func subtractAddressArm64(target, base uintptr) (int64, bool) {
	if target >= base {
		delta := target - base
		if delta > uintptr(^uint64(0)>>1) {
			return 0, false
		}
		return int64(delta), true
	}
	delta := base - target
	if delta > uintptr(^uint64(0)>>1) {
		return 0, false
	}
	return -int64(delta), true
}
