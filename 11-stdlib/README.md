# Go 常用标准库示例

本目录对应基础篇 `140-GoLang教程——标准库精选.md`，运行环境固定为 Go 1.25.1。根目录中的 Go 文件都是独立的 `package main` 示例，需要逐个运行，不能用 `go run .` 一次编译。

## 示例映射

| 文件 | 内容 |
| --- | --- |
| `01_fmt_examples.go` | 格式化输出与格式动词 |
| `02_strings_examples.go` | 字符串查找、分割、连接与构建 |
| `03_strconv_examples.go` | 基本类型解析与格式化 |
| `04_time_examples.go` | 时间点、格式化与时间间隔 |
| `05_json_examples.go` | JSON 编码、解码与结构体标签 |
| `06_os_examples.go` | 文件、目录和环境变量 |
| `07_io_examples.go` | Reader、Writer、复制和读取上限 |
| `08_complete_example.go` | JSON、文件、字符串和时间组合示例 |
| `09_sort_examples.go` | 排序、稳定性与二分查找 |
| `10_rand_examples.go` | 局部伪随机源与可重复序列 |
| `11_filepath_examples.go` | 本机路径拼接、解析和匹配 |
| `12_bytes_examples.go` | 字节切片和内存缓冲区 |
| `13_crypto_rand_examples.go` | 安全随机令牌 |
| `math_demo.go` | 数学、排序、路径、字节和随机数综合演示 |
| `exercises` | 本章 11 道代码练习的参考实现与测试 |

## 运行示例

在代码库根目录执行单个示例：

```bash
go run ./11-stdlib/01_fmt_examples.go
go run ./11-stdlib/13_crypto_rand_examples.go
```

`06_os_examples.go` 和 `08_complete_example.go` 会在当前目录短暂创建文件，正常退出时会清理。建议在具有写权限的工作目录中运行。

## 验证练习

```bash
cd 11-stdlib/exercises
go test ./...
go vet ./...
```

练习测试覆盖文本规范化、数值溢出、时区与夏令时、Timer 超时与 Context 取消、JSON 严格解码、稳定排序、二分查找、可重复随机序列、临时文件和安全随机令牌。
