# Go 测试与质量验证示例

本目录对应基础篇 `150-GoLang教程-测试：写得爽，跑得快.md`，使用 Go 1.25.1。

## 目录内容

- `math.go`、`math_test.go`：最小单元测试、表格驱动测试和基准测试。
- `02-advanced-testing_test.go`：`TestMain`、Helper、并行子测试和最小接口替身。
- `table_driven_test_example_test.go`：独立的表格驱动示例。
- `exercises`：本章 9 道练习的参考实现、测试和基准。

## 验证

```bash
go test -v ./...
go test -race ./...
go test -bench=. -benchmem -count=3
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

cd exercises
go test ./...
go test -race ./...
go test -bench=. -benchmem
go vet ./...
```

`exercises` 中的并发计数器是已经修复的数据竞争版本。练习时可以临时移除 `Counter` 的互斥锁并运行 `go test -race ./...` 观察报告，随后恢复同步；故意存在竞争的实现不作为仓库的最终代码保留。

覆盖率表示语句是否执行，不表示断言是否有效。基准结果依赖 Go 版本、硬件、系统负载和输入；比较实现时应保留多次运行的原始结果，不把单次数据写成普遍性能结论。
