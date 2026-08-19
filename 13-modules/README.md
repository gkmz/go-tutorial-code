# Go 模块管理示例

本目录对应基础篇《工程实践-Go 模块管理》和高级篇《Go 模块管理：依赖图、可复现构建与供应链边界》，统一使用 Go 1.25.1。

## 目录映射

| 目录 | 主题 | 验证方式 |
| --- | --- | --- |
| `01-basic` | `go mod init` 与最小模块 | `go run .` |
| `02-dependencies` | 直接依赖、间接依赖与 `go.sum` | `go run .` |
| `03-replace` | 主模块使用本地 `replace` | `go run .` |
| `04-versions` | 语义化版本与主版本路径 | `go run .` |
| `05-private` | `GOPRIVATE` 与认证配置说明 | `go run .` |
| `06-workspace` | 两个本地模块通过 `go.work` 联调 | `go test ./app/... ./greeter/...`、`go run ./app` |
| `exercises` | 练习 2：本地模块替换的可运行答案 | `go test ./...` |

每个子目录的 `go.mod` 都是独立模块。不要在 `13-modules` 根目录直接运行 `go test ./...`，因为根目录本身不是模块。

## workspace 示例

`06-workspace` 包含 `app` 和 `greeter` 两个模块。`app/go.mod` 只声明模块依赖，不包含本地 `replace`；根目录的 `go.work` 在开发阶段把依赖解析到本地 `greeter`：

```bash
cd 06-workspace
go env GOWORK
go list -m
go test ./app/... ./greeter/...
go run ./app
```

`go list -m` 列出 workspace 中的主模块。这个教学示例让 `app` 以占位版本 `v0.0.0` 要求尚未发布的 `example.com/workspace/greeter`；执行 `go list -m all` 会继续加载完整模块图并尝试查询该未发布版本，因此不作为本示例的本地模块枚举命令。构建、测试和 `go run ./app` 会使用 `go.work` 中的本地 `greeter`。

`06-workspace` 根目录只有 `go.work`，自身不是模块，因此 `go test ./...` 的模式前缀不属于任何 workspace 模块。需要显式使用 `./app/...` 和 `./greeter/...`，或者进入具体模块执行测试。

离开该目录或者设置 `GOWORK=off` 后，`app` 无法仅凭本地源码解析尚未发布的 `example.com/workspace/greeter`。这正是 workspace 影响构建输入的证据，也是 CI 必须明确是否启用 `go.work` 的原因。

## 依赖图与版本诊断

在任意包含依赖的示例模块中，可以运行：

```bash
go list -m all
go mod graph
go mod why -m golang.org/x/text
go list -m -u all
go mod verify
```

`go mod graph` 展示模块需求边，不等于包导入图。`go mod why` 用于解释主模块中的包为什么需要某个包或模块；没有依赖路径时会给出相应说明。

## 验证全部示例

```bash
make test
make run-all
```

`05-private` 只打印配置步骤，不会连接真实私有仓库。代理、认证和私有模块验证必须在组织允许的网络与凭据环境中完成。

文章中的练习 1、3、4、5、6 是依赖环境、命令和凭据策略的操作题，不在代码库中伪造公共仓库、测速结果或私有凭据。练习 2 的完整参考实现位于 `exercises`。
