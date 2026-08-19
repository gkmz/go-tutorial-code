# Go Calculator 项目示例

本目录对应基础篇 `180-GoLang教程——项目实战示例.md`，使用 Go 1.25.1。它演示从模块初始化、核心包、命令行入口、测试、构建到发布说明的完整闭环。

## 目录

- `pkg/calculator`：可测试的计算器核心，返回错误并保护历史记录状态。
- `cmd/calc`：解析参数、选择运算、格式化输出和设置退出码。
- `internal/mathop`：内部包导入边界示例。
- `exercises`：幂运算和开方练习的参考实现与测试。
- `main_urfave.go`：带 `urfave/cli` build tag 的扩展示例，默认不参与普通构建。

## 验证

```bash
gofmt -w ./...
go test ./...
go test -race ./...
go vet ./...
go run ./cmd/calc 10 + 20
go run ./cmd/calc 10 / 0
go run -tags urfave ./cmd/calc --help
cd exercises
go test ./...
go vet ./...
```

CLI 中的 `*` 需要在 Shell 中加引号，例如 `go run ./cmd/calc 5 '*' 3`。项目使用浮点数只为演示计算流程；金额和结算场景应改用最小货币单位整数或经过评估的定点表示。

## 练习映射

练习 1 的幂运算和开方已经放在 `exercises/calculator.go` 并由测试覆盖。交互式模式、配置文件、历史持久化、单位转换和 Cobra 子命令属于扩展项目题，文章给出边界与验收方向；实现时应先补测试，再拆分核心包和入口。
