# Wire 编译期依赖注入示例

本目录对应高级篇《依赖注入与 Wire：从手工装配到编译期代码生成》，使用 Go 1.25.1 和 Wire v0.7.0。

[google/wire](https://github.com/google/wire) 已被官方归档。本示例用于理解和维护已有 Wire 项目，不表示新项目必须采用 Wire。依赖图较小时，`InitializeAppManually` 展示的手工构造函数装配通常更直接。

## 文件职责

| 文件 | 内容 |
| --- | --- |
| `app.go` | Config、Repository、Service、App 和带清理函数的 Provider |
| `manual.go` | 与生成结果等价的手工装配 |
| `wire.go` | 带 `wireinject` 构建标签的 Injector 声明 |
| `wire_gen.go` | Wire 生成的普通 Go 装配代码 |
| `app_test.go` | 生成代码、错误传播和清理行为验证 |
| `exercises` | 手工注入测试替身、多个资源清理顺序等代码练习参考答案 |
| `cmd/demo` | 可运行入口 |

## 重新生成

```bash
cd /Users/hank/workspace/mine/go-tutorial-code/18-advanced
go run github.com/google/wire/cmd/wire@v0.7.0 ./wire
```

生成后检查差异并运行：

```bash
go test ./wire ./wire/cmd/demo
go vet ./wire ./wire/cmd/demo
go test -race ./wire
go run ./wire/cmd/demo
```

练习答案可以单独运行：

```bash
go test ./wire/exercises
```

`wire.go` 只在生成阶段通过 `wireinject` 标签参与分析，普通构建使用已提交的 `wire_gen.go`。部署环境不需要安装 Wire，但 CI 应重新生成并检查 `wire_gen.go` 是否过期。
