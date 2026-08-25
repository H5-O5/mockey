// These pin that making a declared method shadow a promoted one NARROWS
// NOTHING. The worry is real: if the embedded type is unexported and lives in
// another package, a caller cannot write o.inner, so if GetMethod stopped
// handing back the promoted method there might be no way left to reach it.
//
// It stays reachable, through the API whose documented job that is. The two
// entry points now mean different things instead of collapsing onto one:
//   GetMethod       -> the method o.Name() actually dispatches to
//   GetNestedMethod -> the one inside the embedded field
package tests

import (
	"reflect"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/bytedance/mockey/tests/embedpkg"
)

// A method DECLARED on the containing type wins, even when the shadowed one
// comes from an unexported embedded type in another package.
func TestUnexportedEmbeddedFromAnotherPackage(t *testing.T) {
	o := embedpkg.New()
	got := reflect.TypeOf(mockey.GetMethod(o, "Name"))
	if got.In(0) != reflect.TypeOf(embedpkg.Outer{}) {
		t.Fatalf("resolved %v (recv %v), want a method on embedpkg.Outer", got, got.In(0))
	}
	mk := mockey.Mock(mockey.GetMethod(o, "Name")).Return("mocked-outer").Build()
	defer mk.UnPatch()
	if v := o.Name(); v != "mocked-outer" {
		t.Fatalf("mocking the declared method had no effect: %q", v)
	}
}

// Can the shadowed, cross-package, unexported-typed method still be reached?
// GetNestedMethod is the API whose documented job is exactly this.
func TestShadowedStillReachableViaNestedAPI(t *testing.T) {
	o := embedpkg.New()
	m := mockey.GetNestedMethod(o, "Name")
	t.Logf("NESTED GetNestedMethod(*embedpkg.Outer, Name) -> %v", reflect.TypeOf(m))

	mk := mockey.Mock(m).Return("mocked-inner").Build()
	defer mk.UnPatch()

	if got := o.CallInner(); got != "mocked-inner" {
		t.Fatalf("shadowed inner method NOT mockable: got %q", got)
	}
	t.Log("NESTED ok: shadowed method still mockable via GetNestedMethod")

	if got := o.Name(); got != "outer" {
		t.Fatalf("outer method should be untouched, got %q", got)
	}
	t.Log("NESTED ok: outer method untouched")
}
