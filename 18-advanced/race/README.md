# Data Race 检测示例

本目录对应高级篇第 260 章，演示数据竞争的复现、Race Detector 报告和常见修复方式。

正常测试只包含同步正确的实现：

```bash
go test ./race/...
go test -race ./race/...
```

`cmd/racy` 故意包含数据竞争，仅用于观察报告。该命令在检测到竞争后应以非零状态退出：

```bash
go run -race ./race/cmd/racy
```

修复后的对照版本：

```bash
go run -race ./race/cmd/fixed
```

代码练习参考答案位于 `exercises`，覆盖原子化的“检查并扣减”以及使用 channel 建立 happens-before 关系。
