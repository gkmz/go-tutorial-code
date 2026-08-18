# Go 反射实战示例

本目录对应高级篇第 280 章，覆盖 Type、Value、Kind、字段描述与修改、required 标签校验、安全动态调用和性能对比。

```bash
go run ./reflection
go test ./reflection/...
go test -race ./reflection/...
go test -run '^$' -bench . -benchmem ./reflection
```

代码练习参考答案位于 `exercises`，包括按名称复制可赋值字段，以及支持循环指针检测的嵌套 required 校验。
