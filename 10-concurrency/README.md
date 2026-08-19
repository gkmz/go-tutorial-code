# 基础并发示例

本目录对应《GoLang 避坑实战教程：基础篇》的 130 章“并发基础”，使用 Go 1.25.1。

## 文件映射

| 文件 | 内容 |
| --- | --- |
| `01_basic_goroutine.go` | Goroutine 启动和完成同步 |
| `02_channel_basics.go` | 无缓冲、缓冲、关闭和 `range` |
| `03_select.go` | 多路等待、超时、非阻塞接收和取消 |
| `04_waitgroup.go` | `WaitGroup` 与循环变量传参 |
| `05_mutex.go` | 共享状态与数据竞争对比 |
| `06_worker_pool.go` | 固定 worker、任务队列和结果关闭 |
| `07_pitfalls.go` | 常见并发错误写法与修复 |
| `08_context.go` | 超时、手动取消和 goroutine 退出 |
| `09_rate_limiter.go` | 基于 ticker 的限流 |
| `10_fan_in_out.go` | 扇出、扇入和管道 |
| `11_sync_once.go` | `sync.Once` 的一次性初始化 |
| `exercises/` | 练习答案与测试 |

## 运行示例

这些 `.go` 文件是独立的 `main` 示例，需要逐文件运行：

```bash
go run ./01_basic_goroutine.go
go run ./02_channel_basics.go
go run ./06_worker_pool.go
go run ./08_context.go
```

## 验证练习答案

```bash
cd exercises
go test ./...
go test -race ./...
go vet ./...
```

并发示例的输出顺序可能因调度而变化，验证重点是程序能否正常退出、channel 是否正确关闭、取消后 goroutine 是否返回，以及 Race Detector 是否报告数据竞争。
