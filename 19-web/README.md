# Web 篇示例

示例基于 Go 1.25.1。`gin-basic` 和 `graceful` 可以直接运行；`upload-jwt` 展示安全边界，密钥必须通过环境变量提供。

```bash
go run ./gin-basic
JWT_SECRET='replace-with-a-long-random-secret' go run ./upload-jwt
```
