# 高级篇示例

每个子目录都是独立的可执行示例或测试包，使用 Go 1.25.1 验证。

并发进阶练习参考实现位于 `concurrency/exercises`，覆盖受控 BufferPool、互斥锁、读写锁、`sync.Once`、`sync.Map`、atomic、`errgroup.SetLimit`、带权信号量和错误取消传播。运行练习测试：

Context 练习参考实现位于 `context/exercises`，覆盖并发下游调用、取消传播和固定 worker 数量；`context/grpc` 说明截止时间、metadata 和状态码的跨进程边界。

泛型练习参考实现位于 `exercises`，覆盖反转切片、泛型队列、Reduce、Set、GroupBy 和带比较函数的泛型二叉搜索树，并配有测试。

泛型函数和泛型类型的可运行示例位于 `generics`，包括 `Map`、`Filter`、数值约束和不会暴露内部切片的 `Calculator`：

```bash
go run ./generics
go test ./generics
cd exercises && go test ./...
```

其他文章的核心示例目录如下：

| 目录 | 对应主题 |
| --- | --- |
| `gmp` | GMP 调度、抢占、计时器、Trace 与 GOMAXPROCS 实验 |
| `singleflight` | 请求合并、TTL 缓存、独立等待者超时、Forget 与教学实现 |
| `race` | Race Detector 报告、安全修复、故意失败示例与练习答案 |
| `generics` | 泛型约束和泛型函数 |
| `reflection` | Type、Value、字段读写、安全动态调用、练习答案和 benchmark |
| `memory` | GC 内存统计、分配 benchmark、有界缓存练习答案 |
| `escape` | 返回值、返回指针、slice、闭包、编译器诊断和分配 benchmark |
| `pprof` | HTTP pprof、CPU/heap profile 和采集辅助函数 |
| `cgo` | CGO 调用边界、字符串所有权、构建标签和调用 benchmark |
| `config` | 配置优先级、Viper、校验、脱敏摘要和原子快照 |
| `logging` | slog、zap、脱敏、HTTP 请求日志、轮转与 benchmark |
| `fuzz` | 原生模糊测试 |
| `wire` | Wire 生成流程说明 |

这些目录用于支撑教程中的关键示例。依赖注入生成、CGO、pprof 和外部服务相关示例需要根据文章说明准备工具或运行环境。

GMP 示例位于 `gmp`，包含可控数量的 Goroutine 实验、可取消 CPU 任务、计时器等待、不同 `GOMAXPROCS` 的 Benchmark 和 Trace 工作负载：

```bash
go run ./gmp -mode burst -count 10000
go test ./gmp
go test -race ./gmp
go test -run '^TestTraceWorkload$' -trace /tmp/gmp-trace.out ./gmp
go test -run '^$' -bench BenchmarkCPUWorkGOMAXPROCS -benchmem ./gmp
```

singleflight 示例位于 `singleflight`，覆盖稳定的并发重叠测试、缓存组合、key 隔离、等待者超时、错误共享和 `Forget`：

```bash
go run ./singleflight
go test ./singleflight
go test -race ./singleflight
```

Data Race 示例将故意失败的程序与正常测试隔离，避免全量测试被演示代码污染：

```bash
go test -race ./race/...
go run -race ./race/cmd/racy  # 预期报告竞争并返回非零状态
go run -race ./race/cmd/fixed
```

内存管理示例位于 `memory`，可以观察分配、手动 GC 和存活堆变化，也可以运行预分配 benchmark 与有界缓存练习答案：

```bash
go run ./memory
go test ./memory
go test -run '^$' -bench 'Benchmark(GrowingSlice|PreallocatedSlice|BoundedCacheSet)$' -benchmem ./memory
GODEBUG=gctrace=1 go run ./memory
```

逃逸分析示例位于 `escape`，可以分别查看编译器诊断和实际分配次数：

```bash
cd escape
go build -o /tmp/go-escape-example -gcflags='-m=2' .
go test -run '^$' -bench . -benchmem
```

pprof 示例位于 `pprof`，默认只监听 `127.0.0.1:6060`，包含可观察的示例负载、block/mutex profile 配置以及 `runtime/pprof` 文件采集辅助函数：

```bash
go run ./pprof
go test ./pprof
go test -race ./pprof
go tool pprof http://127.0.0.1:6060/debug/pprof/heap
go tool pprof 'http://127.0.0.1:6060/debug/pprof/profile?seconds=10'
```

CGO 示例位于 `cgo`，包含 C 函数调用、Go/C 字符串复制和释放、CGO 构建标签以及 Go/C 调用 benchmark：

```bash
go run ./cgo
go test ./cgo
go test -race ./cgo
go test -run '^$' -bench . -benchmem ./cgo
CGO_ENABLED=0 go run ./cgo
```

配置管理示例位于 `config`，包含文件和环境变量合并、显式覆盖、强类型校验、`.env` 辅助、脱敏摘要和原子快照：

```bash
go test ./config
go vet ./config
go test -race ./config
```

日志管理示例位于 `logging`，包含 slog Handler、敏感字段脱敏、HTTP 请求日志、zap 动态级别、文件轮转和公平 benchmark：

```bash
go test ./logging
go vet ./logging
go test -race ./logging
go test -run '^$' -bench . -benchmem ./logging
```

反射示例位于 `reflection`，覆盖安全字段读取、字段修改、标签校验和性能对比：

```bash
cd reflection
go test ./...
go test -race ./...
go test -run '^$' -bench . -benchmem
```

gRPC 与 Context 的调用链、metadata、拦截器和状态码说明位于 `context/grpc/README.md`。该目录使用 protobuf 生成代码作为类型前提，重点验证 Context 的传播规则。

```bash
go test ./...
go test -race ./...
```
