package inst

import (
	"strings"
	"testing"
	"unsafe"

	"golang.org/x/arch/arm64/arm64asm"
)

func TestArm64HasPCRelative(t *testing.T) {
	if HasPCRelative([]byte{0x1f, 0x20, 0x03, 0xd5}) {
		t.Fatal("NOP must be position independent")
	}
	if !HasPCRelative(encodeWords(t, mustRewriteADR(t, 0x100000, 0x100008, 0))) {
		t.Fatal("ADR must be detected as pc-relative")
	}
	if !HasPCRelative(encodeWords(t, mustRewriteScaled(t, 0x35000000, 0x100000, 0x100040, 19, 0x7ffff, 5))) {
		t.Fatal("CBNZ must be detected as pc-relative")
	}
}

func TestArm64Relocate(t *testing.T) {
	t.Run("ADR becomes move-absolute", func(t *testing.T) {
		oldAddr := uintptr(0x100000)
		newAddr := uintptr(0x200000)
		target := oldAddr + 0x400
		code := encodeWords(t, mustRewriteADR(t, oldAddr, target, 11))

		relocated, err := Relocate(code, oldAddr, newAddr)
		if err != nil {
			t.Fatal(err)
		}
		if len(relocated) != 16 {
			t.Fatalf("relocated length = %d, want 16", len(relocated))
		}
		if got := decodeMovedAddress(t, relocated[:16]); got != target {
			t.Fatalf("relocated ADR target = 0x%x, want 0x%x", got, target)
		}
	})

	t.Run("ADRP becomes move-absolute page", func(t *testing.T) {
		oldAddr := uintptr(0x100000)
		newAddr := uintptr(0x300000)
		target := uintptr(0x34567000)
		base := uint32(0x90000000 | 3)
		enc, err := rewriteADRPImmediate(base, target, oldAddr)
		if err != nil {
			t.Fatal(err)
		}
		relocated, err := Relocate(encodeWords(t, enc), oldAddr, newAddr)
		if err != nil {
			t.Fatal(err)
		}
		if got := decodeMovedAddress(t, relocated[:16]); got != target&^uintptr(0xfff) {
			t.Fatalf("relocated ADRP target = 0x%x, want 0x%x", got, target&^uintptr(0xfff))
		}
	})

	t.Run("CBNZ to copied target widens and remaps", func(t *testing.T) {
		oldAddr := uintptr(0x100000)
		newAddr := uintptr(0x200000)
		enc := mustRewriteScaled(t, 0x35000000, oldAddr, oldAddr+8, 19, 0x7ffff, 5)
		code := encodeWords(t,
			enc,
			0xd503201f, // NOP
			0xd503201f, // NOP
		)

		relocated, err := Relocate(code, oldAddr, newAddr)
		if err != nil {
			t.Fatal(err)
		}
		if len(relocated) != 32 {
			t.Fatalf("relocated length = %d, want 32", len(relocated))
		}
		first, err := arm64asm.Decode(relocated[:4])
		if err != nil {
			t.Fatal(err)
		}
		if first.Op != arm64asm.CBZ {
			t.Fatalf("first relocated instruction = %v, want CBZ", first)
		}
		if got := decodeMovedAddress(t, relocated[4:20]); got != newAddr+28 {
			t.Fatalf("relocated CBNZ target = 0x%x, want 0x%x", got, newAddr+28)
		}
	})

	t.Run("conditional branch widens and inverts the condition", func(t *testing.T) {
		oldAddr := uintptr(0x100000)
		newAddr := uintptr(0x210000)
		base := uint32(0x54000000) // B.EQ
		enc, err := rewriteScaledImmediate(base, oldAddr+8, oldAddr, 19, 0x7ffff, 5)
		if err != nil {
			t.Fatal(err)
		}
		code := encodeWords(t, enc, 0xd503201f, 0xd503201f)

		relocated, err := Relocate(code, oldAddr, newAddr)
		if err != nil {
			t.Fatal(err)
		}
		if got := relocatedWord(relocated[:4]) & 0xf; got != 1 {
			t.Fatalf("inverted condition code = %d, want 1 (NE)", got)
		}
		if got := decodeMovedAddress(t, relocated[4:20]); got != newAddr+28 {
			t.Fatalf("relocated B.cond target = 0x%x, want 0x%x", got, newAddr+28)
		}
	})

	t.Run("LDR literal is rejected loudly", func(t *testing.T) {
		_, err := Relocate([]byte{0x00, 0x00, 0x00, 0x58}, 0x100000, 0x200000)
		if err == nil {
			t.Fatal("expected LDR-literal relocation to be rejected")
		}
		if !strings.Contains(err.Error(), "literal") {
			t.Fatalf("error %q does not mention the literal load", err)
		}
	})
}

func mustRewriteADR(t *testing.T, instAddr, target uintptr, reg uint32) uint32 {
	t.Helper()
	enc, err := rewriteADRImmediate(0x10000000|reg, target, instAddr)
	if err != nil {
		t.Fatal(err)
	}
	return enc
}

func mustRewriteScaled(t *testing.T, base uint32, instAddr, target uintptr, bits uint, mask uint32, shift uint) uint32 {
	t.Helper()
	enc, err := rewriteScaledImmediate(base, target, instAddr, bits, mask, shift)
	if err != nil {
		t.Fatal(err)
	}
	return enc
}

func encodeWords(t *testing.T, words ...uint32) []byte {
	t.Helper()
	out := make([]byte, len(words)*instLen)
	for i, word := range words {
		*(*uint32)(unsafe.Pointer(&out[i*instLen])) = word
	}
	return out
}

func relocatedWord(code []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&code[0]))
}

func decodeMovedAddress(t *testing.T, code []byte) uintptr {
	t.Helper()
	if len(code) < 16 {
		t.Fatalf("move sequence too short: %d", len(code))
	}
	var value uintptr
	for i := 0; i < 4; i++ {
		word := relocatedWord(code[i*instLen:])
		imm := uintptr((word >> 5) & 0xffff)
		shift := uintptr((word >> 21) & 0x3)
		value |= imm << (shift * 16)
	}
	return value
}
