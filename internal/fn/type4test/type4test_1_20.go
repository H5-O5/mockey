//go:build go1.20
// +build go1.20

/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package type4test

var (
	GlobalFn1 = func(int) {}
	GlobalFn2 = func(A0, int) {}
)

func Foo0(int) {}

type A0 struct {
	Inner string
}

func (f A0) Foo(i int)  {}
func (f *A0) Bar(i int) {}

//go:noinline
func Foo[T any](t T) {}

//go:noinline
func NoArgs[T any]() {}

type A[T any] struct {
	Inner T
}

//go:noinline
func (f A[T]) Foo(i int) {}

//go:noinline
func (f *A[T]) Bar(i int, t T) {}

//go:noinline
func (f *A[T]) NoArgs() {}

//go:noinline
func CacheLike[T any](ctx any, client *int, key string, data T, expiration int64) error {
	return nil
}

// stringifyLike mimics utils.JsonMarshal: taking an interface{} forces an
// interface conversion in the generic wrapper, which is the key difference
// from a `return nil` body.
//
//go:noinline
func stringifyLike(v any) string {
	return ""
}

// fakeClient / fakeCmd mimic the *goredis.Client -> *redisv6.IntCmd method
// chain that SetCache invokes, so the generic wrapper has a realistic body.
type FakeCmd struct{}

func (c *FakeCmd) Result() (int64, error) { return 0, nil }

type FakeClient struct{}

// FakeClientInstance is a package-level value the caller package can take its
// address from, mirroring how SetCache is called with a concrete *Client.
var FakeClientInstance = FakeClient{}

func (c *FakeClient) Do(ctx any, key string, val string, exp int64) *FakeCmd { return &FakeCmd{} }

// SetCacheLike mirrors meego_ai's redis.SetCache[string]: a generic function
// whose body performs an interface conversion (data -> any) AND a method call
// chain on a pointer argument. This is the shape that produced the SIGSEGV in
// GenericInfo.Equal on linux/amd64 with -N -l.
//
//go:noinline
func SetCacheLike[T any](ctx any, client *FakeClient, key string, data T, expiration int64) error {
	cacheData := stringifyLike(data)
	_, err := client.Do(ctx, key, cacheData, expiration).Result()
	return err
}

// BigValue is a value (pointer-free) type used as a generic type argument so
// the wrapper emits a convT2E with an explicit type-descriptor LEA, giving the
// wrapper more than one RIP-relative LEA before the body CALL.
type BigValue struct {
	A, B, C int
}

//go:noinline
func SetCacheStruct[T any](ctx any, client *FakeClient, key string, data T, expiration int64) error {
	cacheData := stringifyLike(data)
	_, err := client.Do(ctx, key, cacheData, expiration).Result()
	return err
}
