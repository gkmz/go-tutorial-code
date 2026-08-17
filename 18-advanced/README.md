# 高级篇示例

每个子目录都是独立的可执行示例或测试包，使用 Go 1.25.1 验证。

并发进阶练习参考实现位于 `concurrency/exercises`，覆盖受控 BufferPool、互斥锁、读写锁、`sync.Once`、`sync.Map`、atomic、`errgroup.SetLimit`、带权信号量和错误取消传播。运行练习测试：

Context 练习参考实现位于 `context/exercises`，覆盖并发下游调用、取消传播和固定 worker 数量；`context/grpc` 说明截止时间、metadata 和状态码的跨进程边界。

泛型练习参考实现位于 `exercises`，覆盖反转切片、泛型队列、Reduce、Set、GroupBy 和带比较函数的泛型二叉搜索树，并配有测试。

其他文章的核心示例目录如下：

| 目录 | 对应主题 |
| --- | --- |
| `gmp` | GMP 调度、抢占、计时器、Trace 与 GOMAXPROCS 实验 |
| `singleflight` | 请求合并、TTL 缓存、独立等待者超时、Forget 与教学实现 |
| `race` | Race Detector 测试 |
| `generics` | 泛型约束和泛型函数 |
| `reflection` | Type、Value、字段读取与修改、标签校验和 benchmark |
| `memory` | GC 内存统计 |
| `escape` | 返回值、返回指针、slice、闭包、编译器诊断和分配 benchmark |
| `pprof` | HTTP pprof 入口 |
| `cgo` | CGO 调用边界 |
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

逃逸分析示例位于 `escape`，可以分别查看编译器诊断和实际分配次数：

```bash
cd escape
go build -o /tmp/go-escape-example -gcflags='-m=2' .
go test -run '^$' -bench . -benchmem
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
