# Go 面试现场编码示例

本目录存放《Go 后端面试教程与工程实践》中的可运行代码。每个子目录都是面试题的一种参考实现，均配有测试，可以在本目录执行：

```bash
go test ./...
go test -race ./...
go vet ./...
```

## 目录

| 目录 | 内容 | 对应资料 |
| --- | --- | --- |
| `01-basic-algorithms` | 两数之和、合并区间、滑动窗口、二叉树层序遍历 | 基础数据结构与算法编码题 |
| `02-backend-components` | 可取消批处理、TTL 缓存、分页、事件去重 | Go 后端常见组件编码题 |
| `03-concurrency` | 生产者消费者、并发计数器、请求合并、优雅退出 | 并发编程现场题 |
| `04-lru` | 并发安全 LRU 缓存 | LRU 完整实现与测试 |
| `05-workerpool` | 可取消 Worker Pool | 可取消 Worker Pool 完整实现与测试 |
| `06-httpclient` | HTTP Client、幂等重试与响应体管理 | HTTP Client 封装与重试编码题 |
| `07-ratelimiter` | 可取消令牌桶限流器 | 限流器实现与测试 |
| `08-core-semantics` | Slice、typed nil、defer、错误包装、泛型和反射实验 | Go 语言核心面试 |
| `09-web-examples` | HTTP 超时、请求 ID 和慢消费者广播 | 工程化与 Web 开发 |
| `10-runtime-experiments` | 逃逸、内存快照、CPU Profile、Trace 和 Benchmark | Runtime 与性能优化 |
| `11-storage-examples` | Cache Aside、内存缓存和 Fencing Token | 数据库与存储 |
| `12-rpc-messaging` | RPC Deadline、Kafka 分区、重试和死信 | gRPC 与 Kafka |

这些实现用于学习和面试训练，不是面向所有生产场景的通用库。阅读代码时应同时关注需求约束、失败路径、资源生命周期和测试边界。
