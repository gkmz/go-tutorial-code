# CLI 工具开发示例

使用 `urfave/cli` 框架开发命令行工具的完整示例。

## 目录结构

```
16-cli-urfave/
├── 01-hello/           # 最简单的 Hello World
├── 02-flags/           # 添加全局 Flag
├── 03-commands/        # 添加命令
├── 04-subcommands/     # 子命令示例
├── 05-complete/        # 完整示例
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

## 生命周期

```
Before → Action → After
```

- `Before`：命令执行前调用，可用于验证
- `Action`：命令的主要逻辑
- `After`：命令执行后的收尾钩子；关键资源仍应在创建位置使用 `defer` 管理

## 参考资源

- 官方文档：https://cli.urfave.org/v2/
- GitHub：https://github.com/urfave/cli
