package probe

import (
	"context"
	"reflect"
	"testing"

	"code.byted.org/kv/goredis"
	redis "code.byted.org/kv/redis-v6"
	"github.com/bytedance/mockey"
)

// A goredis.Client whose EMBEDDED *redis.Client is NON-NIL is the field shape:
// a live client built by goredis.NewClient always has it populated. This is the
// path that returns the embedded method and thus the one the fix changes.
func liveClient() *goredis.Client {
	return &goredis.Client{Client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:1"})}
}

func TestRealLiveResolvesOuter(t *testing.T) {
	c := liveClient()
	if c.Client == nil {
		t.Fatal("embedded *redis.Client is nil; not the field shape")
	}
	got := reflect.TypeOf(mockey.GetMethod(c, "WithContext"))
	t.Logf("LIVE GetMethod(*goredis.Client{Client:non-nil}) -> %v", got)
	if want := reflect.TypeOf(&goredis.Client{}); got.Out(0) != want {
		t.Fatalf("resolved %v (out %v), want a method returning %v", got, got.Out(0), want)
	}
}

func TestRealLiveReturnAcceptsOuter(t *testing.T) {
	c := liveClient()
	sentinel := liveClient()
	m := mockey.Mock(mockey.GetMethod(c, "WithContext")).Return(sentinel).Build()
	defer m.UnPatch()
	if got := liveClient().WithContext(context.Background()); got != sentinel {
		t.Fatalf("mock not observed: got %p want %p", got, sentinel)
	}
	t.Log("LIVE acceptance: Return(*goredis.Client) accepted and observed")
}

// REGRESSION GUARD for the real call site in dc/meego_ai:
//   mockey.Mock(mockey.GetMethod(container.Default.Redis.AICache.Client, "Del")).
//       To(func(keys ...string) *redisv6.IntCmd { ... })
//
// goredis does NOT declare Del. It is a purely promoted method reaching
// *goredis.Client through *redis.Client's embedded cmdable, so nothing shadows
// it and GetMethod must still resolve the EMBEDDED one -- the hook there
// returns *redisv6.IntCmd, which only type-checks against the embedded
// signature. This is the case the precedence change could most easily have
// broken, and it is load-bearing for existing user tests.
func TestRealPromotedWithNoDeclaredShadow(t *testing.T) {
	got := reflect.TypeOf(mockey.GetMethod(liveClient(), "Del"))
	t.Logf("PROMOTED GetMethod(*goredis.Client, Del) -> %v", got)
	if want := reflect.TypeOf(&redis.IntCmd{}); got.Out(0) != want {
		t.Fatalf("resolved %v (out %v), want %v", got, got.Out(0), want)
	}
}

func TestRealPromotedDelMockObserved(t *testing.T) {
	sentinel := &redis.IntCmd{}
	m := mockey.Mock(mockey.GetMethod(liveClient(), "Del")).To(func(keys ...string) *redis.IntCmd {
		return sentinel
	}).Build()
	defer m.UnPatch()
	if got := liveClient().Del("k"); got != sentinel {
		t.Fatalf("promoted Del mock not observed: got %p want %p", got, sentinel)
	}
	t.Log("PROMOTED Del: mock observed through *goredis.Client")
}
