# Go 原生模糊测试示例

本目录对应高级篇第 37、38 章，使用 Go 1.25.1。

## 示例映射

| 文件 | 内容 |
| --- | --- |
| `reverse.go` | 字节反转、UTF-8 校验和按码点反转 |
| `reverse_test.go` | 可逆性、UTF-8 有效性和非法输入属性 |
| `record.go` | 带输入大小限制和业务校验的 JSON 解析器 |
| `record_test.go` | 成功解析后的编码、再解析往返属性 |
| `testdata/fuzz/FuzzReverseUTF8` | 已固定到仓库的失败输入回归语料 |

## 常规回归

普通 `go test` 会执行单元测试、所有 fuzz target 的种子语料，以及 `testdata/fuzz` 中已经保存的失败输入：

```bash
cd /Users/hank/workspace/mine/go-tutorial-code/18-advanced
go test ./fuzz
go test -race ./fuzz
```

## 主动模糊测试

一次命令只能主动 fuzz 一个匹配目标：

```bash
go test ./fuzz -fuzz '^FuzzReverseUTF8$' -fuzztime=10s
go test ./fuzz -fuzz '^FuzzRecordRoundTrip$' -fuzztime=10s
```

主动 fuzz 会使用 CPU、内存和本地 fuzz cache。持续时间、并行 worker 和输入资源上限应根据 CI 容量设置。生成过程中发现但尚未失败的有趣输入通常保存在构建缓存中；需要长期回归的失败输入应审查后提交到 `testdata/fuzz/<FuzzName>/`。
