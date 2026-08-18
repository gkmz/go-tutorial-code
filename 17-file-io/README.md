# 文件与 IO 示例

本目录对应高级篇《文件与 IO 操作》，使用 Go 1.25.1 验证。

根模块包含 Reader/Writer、大小限制、Scanner、缓冲写入和流式复制示例：

```bash
go run .
go test ./...
go vet ./...
go test -race ./...
```

`exercises` 是独立模块，包含流式文件复制、ASCII 小写转换与双目标写入、词频统计的参考答案和测试：

```bash
cd exercises
go test ./...
go vet ./...
go test -race ./...
```

示例只在临时目录中创建测试文件。处理用户路径、网络响应和压缩文件时，还需要按教程设置路径、总量、超时和展开大小边界。
