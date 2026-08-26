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
