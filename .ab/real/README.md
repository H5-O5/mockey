# 真实依赖验收（linux/amd64）

用 `dc/meego_ai` 真实 resolved 的依赖图，跑真实的
`code.byted.org/kv/goredis@v5.7.3` + `code.byted.org/kv/redis-v6@v1.1.5`，
不用任何本地替身。`go.mod` / `go.sum` 直接抄自该仓，只加一行
`replace github.com/bytedance/mockey => /mockey`。

需要 go1.24（该仓 `go.mod` 要求）与宿主机已填充的 module cache
（内网模块 docker 里下不到，所以只读挂载 `GOMODCACHE`）。

```sh
docker run --rm --platform linux/amd64 \
  -v "$PWD":/mockey -v "$PWD/.ab/real":/w -w /w \
  -v ~/go/pkg/mod:/gomod:ro \
  -e GOPATH=/gopath -e GOMODCACHE=/gomod -e GOFLAGS=-mod=mod \
  -e GOPROXY=off -e GOSUMDB=off -e GOTOOLCHAIN=local -e MOCKEY_CHECK_GCFLAGS=false \
  golang:1.24-bookworm go test ./probe/... -gcflags=all=-l -count=1 -v
```

## 观察到的结果

修复在位：`-gcflags=all="-N -l"` 与 `=-l` 两种配置全部 ALL OK。

负向控制（`git show 50a7268^:utils.go > utils.go`，真正回退）在真实类型上
精确复现现场 panic，与报告逐字一致：

```
resolved func(*redis.Client, context.Context) *redis.Client (out *redis.Client),
  want a method returning *goredis.Client
panic: return args not match: target: func(*redis.Client, context.Context)
  *redis.Client, index: 0, current type: *goredis.Client
```

⚠️ 注意 `git stash push -- utils.go` 在该文件已与 HEAD 一致时是 **no-op**，
用它做负向控制会得到假阴性（测的其实还是修复版）。要用 `git show <commit>^:utils.go`。

## 为什么必须是 non-nil 的内嵌字段

`goredis.Client` 的 `WithContext` 是指针接收者，`redis.Client` 的也是。
只有当内嵌的 `*redis.Client` **非 nil**（`goredis.NewClient` 造出来的活客户端
总是如此）时，`getMethod` 的匿名字段递归才会真的返回内嵌方法，
从而走到被修复的那条路径。typed-nil 的 `*goredis.Client` 走的是
`3b9352a` 已经修好的 deferred 分支，因此**测不出这个 bug**。

trace 实证（内嵌搜索确实先返回了内嵌方法）：

```
TRACE enter type=*goredis.Client nested=false
TRACE enter type=*redis.Client   nested=true
TRACE   enter/leave redis.baseClient  ok=false
TRACE   enter/leave redis.cmdable     ok=false
TRACE leave type=*redis.Client   ok=true ret=func(*redis.Client, context.Context) *redis.Client
TRACE leave type=*goredis.Client ok=true ret=func(*goredis.Client, context.Context) *goredis.Client
```
