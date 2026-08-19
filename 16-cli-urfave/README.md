# CLI 工具开发示例

本目录对应基础篇《使用 CLI 开发命令行程序》，使用 Go 1.25.1 和 [urfave/cli v2](https://cli.urfave.org/v2/)。示例按复杂度递进，练习模块使用缓冲区测试完整命令行为。

## 目录结构

```
16-cli-urfave/
├── 01-hello/           # 最简单的 Hello World
├── 02-flags/           # 添加全局 Flag
├── 03-commands/        # 添加命令
├── 04-subcommands/     # 子命令示例
├── 05-complete/        # 完整示例
├── exercises/           # 可测试的 CLI 练习答案
└── README.md
```

## 快速开始

```bash
# 进入示例目录
cd 01-hello

# 安装依赖
go mod tidy

# 运行
go run main.go

# 构建
go build -o my-cli main.go
./my-cli
```

## 核心概念

### App
应用的主体，包含名称、版本、命令等信息。

### Command
命令，可以嵌套子命令。

### Flag
选项/参数，分为全局 Flag 和命令 Flag。

### Action
命令的执行逻辑。

### Context
运行时上下文，用于获取 Flag 值、访问 App 信息等。

## 常用 Flag 类型

- `StringFlag`：字符串
- `IntFlag`：整数
- `BoolFlag`：布尔值
- `StringSliceFlag`：字符串数组
- `DurationFlag`：时间间隔

本示例保留框架默认的 `-v/--version`，因此详细模式使用 `-V/--verbose`，不要为 verbose 再注册 `-v`。
同样地，`h` 通常已被内置 `help` 命令占用，业务命令不要重复注册该别名。

## 生命周期

```
Before → Action → After
```

- `Before`：命令执行前调用，可用于验证
- `Action`：命令的主要逻辑
- `After`：命令执行后的收尾钩子；关键资源仍应在创建位置使用 `defer` 管理

## 参考资源

- 官方文档：[cli.urfave.org/v2](https://cli.urfave.org/v2/)
- 源码仓库：[github.com/urfave/cli](https://github.com/urfave/cli)

## 验证

每个示例都是独立模块：

```bash
for dir in 01-hello 02-flags 03-commands 04-subcommands 05-complete; do
  (cd "$dir" && go test ./... && go vet ./...)
done

cd exercises
go test ./...
go test -race ./...
go vet ./...
```

练习答案覆盖 `greet hello --name`、`hello --upper` 和 `count --number` 的完整参数解析、输出捕获与负数校验。多平台构建仍通过 `GOOS`、`GOARCH` 练习；示例中的天气和当前时间输出依赖运行时，不应做精确字符串断言。

练习模块的公开入口是 `exercises.NewApp`，错误路径使用 `ErrNegativeCount` 判断。测试通过 `cli.App.Run` 在进程内模拟真实命令行，不调用 `os.Exit`。
