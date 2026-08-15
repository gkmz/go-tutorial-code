# Wire 示例

Wire 是编译期依赖注入工具。完整项目需要先安装生成器：

```bash
go install github.com/google/wire/cmd/wire@latest
wire
```

`wire.go` 使用 `//go:build wireinject`，生成的 `wire_gen.go` 才是正常构建使用的代码。生成文件应提交到仓库，避免部署环境依赖 Wire 工具。
