//go:build go1.18

package fn

import (
	"reflect"
	"testing"
	"unsafe"
)

func genericDictionary(entries ...reflect.Type) (GenericInfo, []uintptr) {
	words := make([]uintptr, len(entries))
	for i, typ := range entries {
		words[i] = (*[2]uintptr)(unsafe.Pointer(&typ))[1]
	}
	return GenericInfo(uintptr(unsafe.Pointer(&words[0]))), words
}

func TestGenericInfoEqualEquivalentDictionaries(t *testing.T) {
	left, keepLeft := genericDictionary(reflect.TypeOf(""))
	right, keepRight := genericDictionary(reflect.TypeOf(""))
	if left == right {
		t.Fatal("test requires distinct dictionary addresses")
	}
	if !left.Equal(right) {
		t.Fatal("dictionaries for the same concrete type should match")
	}
	_ = keepLeft
	_ = keepRight
}

func TestGenericInfoEqualDifferentDictionaries(t *testing.T) {
	left, keepLeft := genericDictionary(reflect.TypeOf(""))
	right, keepRight := genericDictionary(reflect.TypeOf(int(0)))
	if left.Equal(right) {
		t.Fatal("dictionaries for different concrete types should not match")
	}
	_ = keepLeft
	_ = keepRight
}

// TestGenericInfoEqualInvalidAddressDoesNotFault guards the dereference that
// 77049ca introduced. inst.GetGenericAddr can, for wrappers it does not fully
// model, hand back a non-nil address that is not a readable dictionary. The
// remote failure (dc/meego_ai dal_test, SIGSEGV in GenericInfo.Equal at
// generic.go:57) was exactly this: a high unmapped address such as
// 0x117d6bcf0. Equal must report "not equal" (the mock then falls through to
// the original function) rather than dereference and fault the whole process.
func TestGenericInfoEqualInvalidAddressDoesNotFault(t *testing.T) {
	good, keep := genericDictionary(reflect.TypeOf(""))
	_ = keep

	cases := map[string]uintptr{
		"low-null-page":  0x1,
		"page-aligned-1": 0x1000,
		"high-unmapped":  0x117d6bcf0, // the exact address from the remote SIGSEGV
	}
	for name, bad := range cases {
		t.Run(name, func(t *testing.T) {
			// Both directions must be safe: the bad address can arrive as the
			// analyzer's target dictionary or as the runtime dictionary.
			if GenericInfo(bad).Equal(good) {
				t.Fatalf("invalid GenericInfo 0x%x must not equal a real dictionary", bad)
			}
			if good.Equal(GenericInfo(bad)) {
				t.Fatalf("real dictionary must not equal invalid GenericInfo 0x%x", bad)
			}
			// Two identical invalid values hit the pointer-identity fast path
			// and must report equal without probing (no fault).
		})
	}
}

func TestDerefUintptrSafe(t *testing.T) {
	x := uintptr(0xdeadbeef)
	got, ok := derefUintptrSafe(uintptr(unsafe.Pointer(&x)))
	if !ok || got != x {
		t.Fatalf("readable pointer: got (%#x, %v), want (%#x, true)", got, ok, x)
	}
	if _, ok := derefUintptrSafe(0x1); ok {
		t.Fatal("unmapped address must report not-readable")
	}
}
