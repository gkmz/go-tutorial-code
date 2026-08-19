# Go 项目结构说明

本目录对应基础篇 `170-工程实践-项目结构.md`，使用 Go 1.25.1。文章的重点是目录边界、依赖方向和按规模演进，不是提供一个必须照搬的目录模板。

## 目录说明

| 概念 | 作用 |
| --- | --- |
| `cmd/` | 每个可执行程序的入口 |
| `internal/` | 受 Go 编译器导入规则保护的内部实现 |
| `pkg/` | 明确承诺可供外部使用的稳定库代码，可选 |
| `api/` | OpenAPI、Protobuf 或 GraphQL 等接口定义 |
| `configs/` | 不含密钥的配置模板 |
| `testdata/` | 测试输入和固定样本，按 Go 约定可由测试使用 |

`cmd`、`internal`、`pkg` 等是常见约定，不是 Go 官方强制的项目布局。小项目可以从一个模块和少数文件开始，只有当入口、业务边界或团队协作确实需要时再拆包。

## 可运行示例

`exercises` 包含两个可执行入口：

- `cmd/hello` 调用 `internal/message.Format`；
- `cmd/goodbye` 调用 `internal/message.Farewell`；
- `internal/message` 保存两个入口共享的业务逻辑和测试。

```bash
cd exercises
go test ./...
go vet ./...
go build ./...
go run ./cmd/hello Ada
go run ./cmd/goodbye Ada
```

## 文章练习

本章练习是结构和命令操作题，不需要仓库维护一套虚构的业务实现。完成练习时可以为每个练习创建临时目录，并执行：

```bash
go mod init example.com/layout-practice
go test ./...
go build ./...
```

练习 2 的参考答案是两个入口复用同一个 `internal` 包。练习 3 应观察从允许范围外导入 `internal` 时的编译错误；故意无法编译的代码不保留在答案模块中。其余开放式重构练习应记录拆包原因、依赖方向和验证结果。

## 验收标准

- 入口只负责配置解析、依赖装配、启动和退出；
- 业务逻辑位于可测试的包，不依赖 `main`；
- 包依赖方向清晰，没有 import cycle；
- 敏感配置、构建产物和本地数据不进入版本控制；
- README 说明每个非显然目录的职责和验证命令。
