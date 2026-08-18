# 日志管理示例

本目录对应高级篇《工程实践-日志管理：从事件记录到集中检索》，使用 Go 1.25.1 验证。

示例覆盖：

- 标准库 `log` 多输出；
- slog JSON Handler、固定上下文和递归敏感字段过滤；
- 保留可选 `http.ResponseWriter` 接口的请求日志中间件；
- zap 动态级别、受 Bearer Token 保护的管理 Handler；
- lumberjack 文件轮转；
- `log`、slog 和 zap 的同条件 benchmark。

运行验证：

```bash
go test ./logging
go vet ./logging
go test -race ./logging
go test -run '^$' -bench . -benchmem ./logging
```

Bearer Token 示例只演示最小鉴权边界。生产管理端还需要 TLS、权限控制、审计、限流、密钥管理和级别自动恢复。
