# Go 语言完整教程 - 配套代码

[![Go Version](https://img.shields.io/badge/Go-1.25.1-blue.svg)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Tutorial](https://img.shields.io/badge/Tutorial-hankmo.com-orange.svg)](https://hankmo.com)

## 📚 教程结构

本代码库配套《极客老墨 GoLang 避坑实战教程》，示例按基础、工程实践、高级和 Web 主题组织。教程文章中的完整代码以本仓库目录为准。

### 第一部分：基础教程

- **目标**：让零基础读者能写出简单的 Go 程序
- **状态**: ✅基础示例已完成

| 目录 | 内容 |
| --- | --- |
| `00-setup` | 环境准备 |
| `01-intro` | Go 简介与程序结构 |
| `02-style` | 代码风格 |
| `03-var` | 变量、类型、指针和作用域 |
| `04-const` | 常量、`iota` 和枚举模式 |
| `05-func` | 函数、闭包、方法和 `defer` |
| `06-control` | 控制结构 |
| `07-collections` | 数组、切片和 Map |
| `08-struct-interface` | 结构体、接口和依赖注入 |
| `09-error-handling` | 错误处理和重试 |
| `10-concurrency` | Goroutine、Channel 和并发工具 |
| `11-stdlib` | 常用标准库 |
| `12-testing` | 单元测试、表格测试和测试辅助工具 |

### 第二部分：工程实践和高级教程

- **目标**：深入底层原理，掌握高级编程技巧
- **状态**: 基础工程示例完整，高级主题提供核心示例

| 目录 | 内容 | 状态 |
| --- | --- | --- |
| `13-modules` | 多模块、依赖、`replace` 和版本 | ✅ |
| `14-project-layout` | 项目目录结构说明 | ✅ 文档示例 |
| `15-project-example` | 计算器项目、测试和构建 | ✅ |
| `16-cli-urfave` | urfave/cli 命令行程序 | ✅ |
| `17-file-io` | 文件、流和 IO 组合 | ✅ |
| `18-advanced/concurrency` | worker pool 和取消 | ✅ 核心示例 |
| `18-advanced/context` | Context 超时和取消 | ✅ 核心示例 |
| `18-advanced/generics` | 泛型约束和泛型函数 | ✅ 核心示例 |
| `18-advanced/gmp` | 调度器运行时信息 | ✅ 基础示例 |
| `18-advanced/memory` | 内存统计 | ✅ 基础示例 |
| `18-advanced/pprof` | pprof HTTP 入口 | ✅ 基础示例 |
| `18-advanced/race` | 竞态检测测试 | ✅ |
| `18-advanced/fuzz` | 原生模糊测试 | ✅ |

`18-advanced` 当前是高级教程的关键代码示例，不代表已经覆盖教程中的每一个高级主题；CGO、Wire、singleflight 等内容仍以教程文章和后续补充为准。

### 第三部分：Web 开发实战

- **目标**：从零到一构建企业级 Web 应用
- **状态**: ✅核心示例已补齐，综合项目持续完善

高级篇示例位于 `18-advanced/`，Web 篇示例位于 `19-web/`。每个子目录都是独立的 Go 包或模块，统一使用 Go 1.25.1。

| 目录 | 内容 | 状态 |
| --- | --- | --- |
| `19-web/gin-basic` | Gin 路由和健康检查 | ✅ |
| `19-web/gorm-crud` | GORM 模型、迁移和 CRUD | ✅ |
| `19-web/upload-jwt` | 安全上传和 JWT 算法校验 | ✅ 核心示例 |
| `19-web/graceful` | HTTP 服务优雅关闭 | ✅ |

当前代码库没有完整博客 API、Wire 生成项目、Prometheus 部署编排等综合工程；这些属于教程中的扩展内容，不能直接视为已完成的生产项目。

## 🚀 快速开始

### 环境要求

- Go 1.25.1
- Git

### 克隆代码

```bash
git clone https://github.com/hankmo/go-tutorial-code.git
cd go-tutorial-code
```

### 运行示例

基础篇的大多数 `.go` 文件是独立示例，应按文件运行；带 `go.mod` 的目录是独立模块，应在模块目录中测试。

```bash
# 运行单文件示例
go run ./01-intro/hello.go

# 运行项目示例
cd 15-project-example
go run ./cmd/calc 10 + 20

# 运行 CLI 完整示例
cd ../16-cli-urfave/05-complete
go run . --help

# 运行高级或 Web 示例
cd ../../18-advanced
go run ./generics
cd ../19-web
go run ./gorm-crud
```

### 全量验证

```bash
./scripts/verify.sh
```

验证脚本会执行所有 `go.mod` 模块的测试，并逐个运行无外部依赖的基础单文件示例。它不会启动长期运行的 HTTP 服务，也不会执行需要真实私有仓库、数据库或外部服务的示例。

## 📖 配套教程

完整教程请访问：[https://hankmo.com](https://hankmo.com)

## 🗂️ 目录结构

```
go-tutorial-code/
├── 00-setup/           # 环境准备
├── 01-intro/           # Go 简介
├── 03-var/             # 变量和类型
├── 09-error-handling/  # 错误处理
├── 10-concurrency/    # 基础并发
├── 11-stdlib/         # 标准库
├── 12-testing/        # 测试
├── 13-modules/        # 模块管理
├── 15-project-example/ # 项目实战
├── 16-cli-urfave/     # CLI 实战
├── 17-file-io/        # 文件和 IO
├── 18-advanced/       # 高级核心示例
├── 19-web/            # Web 核心示例
├── common/            # 公共配置
├── docs/              # 图片等文档资源
└── scripts/           # 验证脚本
```

## 💡 代码规范

所有代码遵循以下规范：

- ✅ 使用 Go 1.25.1 版本
- ✅ 通过 `gofmt`、`go test` 和 `go vet` 验证
- ✅ 并发示例支持 `go test -race`
- ✅ 测试代码使用表格测试和测试辅助函数
- ✅ 关键公开符号提供中文 GoDoc
- ✅ 高级篇和 Web 篇模块提供 README 和独立运行说明

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

## 👨‍💻 关于作者

**极客老墨** - 一个热爱折腾的开发者

- 博客：[https://hankmo.com](https://hankmo.com)
- GitHub：[@hankmo](https://github.com/hankmor)
- 公众号：极客老墨

<img src="./docs/weixinqr.jpg" width="200">

---

**极客老墨，继续折腾！** 💪
