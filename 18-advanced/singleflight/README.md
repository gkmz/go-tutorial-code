# singleflight 请求合并示例

本目录对应高级篇 240 和 250 章，使用 Go 1.25.1 与 `golang.org/x/sync v0.22.0` 验证。

## 用法篇

`ArticleService` 展示以下工程边界：

- 使用长期复用的 `singleflight.Group` 合并同 key 调用；
- key 包含租户、语言和文章 ID；
- 缓存未命中后进入共享函数，并执行第二次缓存检查；
- 每个等待者可以通过自己的 Context 提前返回；
- 共享加载使用独立超时，不由首个短超时等待者直接取消；
- 成功结果写入 TTL 缓存，错误默认不缓存；
- `Forget` 只允许后续调用启动，不取消旧调用。

运行示例和测试：

```bash
go run ./singleflight
go test ./singleflight
go test -race ./singleflight
```

示例中的 `context.WithoutCancel` 会保留首个调用 Context 的 Value。真实服务需要明确追踪、鉴权和租户信息如何进入共享加载，不能把请求 Context 的所有权问题交给 singleflight 自动决定。
