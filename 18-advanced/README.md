# 高级篇示例

每个子目录都是独立的可执行示例或测试包，使用 Go 1.25.1 验证。

并发进阶练习参考实现位于 `concurrency/exercises`，覆盖互斥锁、读写锁、`sync.Once`、`sync.Map`、`errgroup.SetLimit`、带权信号量和错误取消传播。运行练习测试：

Context 练习参考实现位于 `context/exercises`，覆盖并发下游调用、取消传播和固定 worker 数量。

泛型练习参考实现位于 `exercises`，覆盖反转切片、泛型队列、Reduce、Set、GroupBy 和带比较函数的泛型二叉搜索树，并配有测试。

其他文章的核心示例目录如下：

| 目录 | 对应主题 |
| --- | --- |
| `gmp` | GMP 调度实验 |
| `singleflight` | 请求合并 |
| `race` | Race Detector 测试 |
| `generics` | 泛型约束和泛型函数 |
| `reflection` | Type、Value 和字段读取 |
| `memory` | GC 内存统计 |
| `escape` | 逃逸分析入口 |
| `pprof` | HTTP pprof 入口 |
| `cgo` | CGO 调用边界 |
| `fuzz` | 原生模糊测试 |
| `wire` | Wire 生成流程说明 |

这些目录用于支撑教程中的关键示例。依赖注入生成、CGO、pprof 和外部服务相关示例需要根据文章说明准备工具或运行环境。

gRPC 与 Context 的调用链、metadata、拦截器和状态码说明位于 `context/grpc/README.md`。该目录使用 protobuf 生成代码作为类型前提，重点验证 Context 的传播规则。

```bash
go test ./...
go test -race ./...
```
