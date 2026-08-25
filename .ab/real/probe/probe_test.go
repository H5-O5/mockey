package probe

import (
	"context"
	"reflect"
	"testing"

	"code.byted.org/kv/goredis"
	redis "code.byted.org/kv/redis-v6"
	"github.com/bytedance/mockey"
)

// These exercise the REAL types from the field report:
//   code.byted.org/kv/goredis.Client embeds *code.byted.org/kv/redis-v6.Client
//   and declares its own WithContext returning *goredis.Client.
// No local stand-ins.

func TestRealTypeShape(t *testing.T) {
	gt := reflect.TypeOf(goredis.Client{})
	f, ok := gt.FieldByName("Client")
	if !ok || !f.Anonymous {
		t.Fatalf("goredis.Client no longer embeds Client anonymously: %+v", f)
	}
	if want := reflect.TypeOf(&redis.Client{}); f.Type != want {
		t.Fatalf("embedded field is %v, want %v", f.Type, want)
	}
	gm, ok := reflect.PtrTo(gt).MethodByName("WithContext")
	if !ok {
		t.Fatal("*goredis.Client has no WithContext")
	}
	rm, ok := reflect.PtrTo(reflect.TypeOf(redis.Client{})).MethodByName("WithContext")
	if !ok {
		t.Fatal("*redis.Client has no WithContext")
	}
	t.Logf("REAL *goredis.Client.WithContext -> %v", gm.Type.Out(0))
	t.Logf("REAL *redis.Client.WithContext   -> %v", rm.Type.Out(0))
	if gm.Type.Out(0) == rm.Type.Out(0) {
		t.Fatal("expected the two WithContext to differ in return type")
	}
}

// TestRealGetMethodResolvesOuter is the requirement: GetMethod must resolve the
// method `c.WithContext(...)` actually dispatches to, i.e. goredis's own.
func TestRealGetMethodResolvesOuter(t *testing.T) {
	var c *goredis.Client // typed nil is enough: resolution is type-directed
	got := reflect.TypeOf(mockey.GetMethod(c, "WithContext"))
	t.Logf("REAL GetMethod(*goredis.Client, WithContext) -> %v", got)

	want := reflect.TypeOf(&goredis.Client{})
	if got.Out(0) != want {
		t.Fatalf("resolved %v (out %v), want a method returning %v", got, got.Out(0), want)
	}
	if got.In(0) != want {
		t.Fatalf("resolved receiver %v, want %v", got.In(0), want)
	}
}

// TestRealEmbeddedStillResolvable pins that the embedded type keeps resolving
// to its own method: the fix must not make redis.Client unmockable.
func TestRealEmbeddedStillResolvable(t *testing.T) {
	var rc *redis.Client
	got := reflect.TypeOf(mockey.GetMethod(rc, "WithContext"))
	t.Logf("REAL GetMethod(*redis.Client, WithContext) -> %v", got)
	if want := reflect.TypeOf(&redis.Client{}); got.Out(0) != want {
		t.Fatalf("resolved %v (out %v), want %v", got, got.Out(0), want)
	}
}

// TestRealReturnAcceptsOuter is the acceptance path: this is the exact
// Mock(GetMethod(...)).Return(...) shape that produced
//   "return args not match: target: func(*redis.Client, context.Context)
//    *redis.Client, index: 0, current type: *goredis.Client"
// in the field. It must build AND the mocked call must observe the mock.
func TestRealReturnAcceptsOuter(t *testing.T) {
	var c *goredis.Client
	sentinel := &goredis.Client{}

	m := mockey.Mock(mockey.GetMethod(c, "WithContext")).Return(sentinel).Build()
	defer m.UnPatch()

	live := &goredis.Client{}
	if got := live.WithContext(context.Background()); got != sentinel {
		t.Fatalf("mock not observed: got %p, want %p", got, sentinel)
	}
	t.Log("REAL acceptance: Return(*goredis.Client) accepted and observed")
}
