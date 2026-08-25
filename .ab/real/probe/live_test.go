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
