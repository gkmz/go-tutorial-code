# 高级篇示例

每个子目录都是独立的可执行示例或测试包，使用 Go 1.25.1 验证。

并发进阶练习参考实现位于 `concurrency/exercises`，覆盖读写锁、`sync.Once`、`sync.Map`、并发限制和 `errgroup` 错误传播。运行练习测试：

```bash
go test ./...
go test -race ./...
```
