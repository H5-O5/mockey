// Package embedpkg models the one shape where the precedence change could
// genuinely REMOVE capability: an unexported embedded type from ANOTHER
// package. A caller outside this package cannot name `inner` or write
// `o.inner`, so if GetMethod stops handing back the promoted method there may
// be no other way to reach it.
package embedpkg

type inner struct{}

//go:noinline
func (inner) Name() string { return "inner" }

// Outer declares Name itself, shadowing inner's.
type Outer struct {
	inner
}

//go:noinline
func (Outer) Name() string { return "outer" }

func New() *Outer { return &Outer{} }

// CallInner reaches the shadowed method from inside the package, so a test can
// tell whether mocking it had any effect.
func (o *Outer) CallInner() string { return o.inner.Name() }
