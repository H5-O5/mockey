//go:build go1.20

package type4caller

import "github.com/bytedance/mockey/internal/fn/type4test"

func CallCacheLikeString(ctx any, client *int, key, data string, expiration int64) error {
	return type4test.CacheLike[string](ctx, client, key, data, expiration)
}

func CallSetCacheLikeString(ctx any, key, data string, expiration int64) error {
	return type4test.SetCacheLike[string](ctx, &type4test.FakeClientInstance, key, data, expiration)
}

func CallSetCacheStruct(ctx any, data type4test.BigValue, expiration int64) error {
	return type4test.SetCacheStruct[type4test.BigValue](ctx, &type4test.FakeClientInstance, "key", data, expiration)
}
