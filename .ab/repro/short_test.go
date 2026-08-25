package repro

import (
	"context"
	"reflect"
	"runtime"
	"testing"

	"github.com/bytedance/mockey"
)

// A SHORT inner method: the optimizer can compress this below the long-jump
// entry sequence, making it unpatchable by UPSTREAM mockey.
type ShortInner struct{ n string }

func (c *ShortInner) WithContext(ctx context.Context) *ShortInner { return c }

type ShortOuter struct {
	*ShortInner
}

func (c *ShortOuter) WithContext(ctx context.Context) *ShortOuter {
	c.ShortInner = c.ShortInner.WithContext(ctx)
	return c
}

func TestShortFuncSize(t *testing.T) {
	m, _ := reflect.TypeOf(&ShortInner{}).MethodByName("WithContext")
	f := runtime.FuncForPC(m.Func.Pointer())
	t.Logf("SIZE inner WithContext entry=%#x name=%s", m.Func.Pointer(), f.Name())
	got := reflect.TypeOf(mockey.GetMethod(&ShortOuter{ShortInner: &ShortInner{}}, "WithContext"))
	t.Logf("SIZE GetMethod -> %v", got)
}

func TestShortE2E(t *testing.T) {
	c := &ShortOuter{ShortInner: &ShortInner{n: "x"}}
	m := mockey.Mock(mockey.GetMethod(c, "WithContext")).Return(c).Build()
	defer m.UnPatch()
	_ = c.WithContext(context.Background())
	t.Log("SHORT E2E called ok")
}
