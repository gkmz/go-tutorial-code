# GMP 调度器实验

本目录对应高级篇 230 章，使用 Go 1.25.1 验证。实验用于观察调度现象，不提供跨机器通用的性能结论。

## Goroutine 数量与内存

程序使用屏障让所有 Goroutine 同时停在接收操作上，再读取 Goroutine 数量和堆统计。默认数量为 10000，可以逐步调整：

```bash
go run ./gmp -mode burst -count 1000
go run ./gmp -mode burst -count 10000
go run ./gmp -mode burst -count 100000
```

不要直接把数量提高到一百万。先观察机器可用内存、运行时间和系统负载。

## 抢占与业务取消

该实验把 `GOMAXPROCS` 临时设置为 1。CPU 任务在计算块内部没有显式让出操作，观察任务依靠运行时调度获得执行机会；CPU 任务仍然通过 Context 检查主动结束。

```bash
go run ./gmp -mode preempt -duration 500ms
```

`observer_ticks` 大于 0 只能说明观察任务获得了执行机会，不能当作精确的抢占延迟指标。

该模式会调用 `runtime.GOMAXPROCS` 并在结束前恢复原值。Go 1.25 中，任何手动调用都会关闭当前进程的自动 `GOMAXPROCS` 更新，因此这个实现只适合短命实验程序，不应原样放入长期运行服务。

## 计时器等待

```bash
go run ./gmp -mode timer -duration 100ms
```

主 Goroutine 等待计时器时不占用 P，其他可运行任务仍可继续执行。

## 测试、Benchmark 与 Trace

```bash
go test ./gmp
go test -race ./gmp
go test -run '^TestTraceWorkload$' -trace /tmp/gmp-trace.out ./gmp
go tool trace /tmp/gmp-trace.out
go test -run '^$' -bench BenchmarkCPUWorkGOMAXPROCS -benchmem ./gmp
```

Benchmark 会改变进程内的 `GOMAXPROCS`，并在每个子测试结束后恢复原值。不同机器的 CPU、调频、后台负载和容器配额不同，不应直接比较绝对数字。

## 练习答案

`exercises.go` 提供一个可取消的有界任务执行器，回答教程中“为 CPU 任务增加并发上限”的练习。测试会记录实际并发数，验证它不会超过配置上限，并检查 Context 取消能够停止继续分发任务：

```bash
go test -run '^TestRunLimited' ./gmp
go test -race -run '^TestRunLimited' ./gmp
```
