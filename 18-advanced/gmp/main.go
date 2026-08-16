package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// GOMAXPROCS 返回当前运行时允许并行执行 Go 代码的 P 数量。
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// 用 WaitGroup 等待一百万个短任务完成，观察 Goroutine 的创建和调度成本。
	const count = 1_000_000
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 任务只做极少量工作，避免把示例变成计算性能测试。
			_ = 1 + 1
		}()
	}
	wg.Wait()

	fmt.Printf("Finished %d goroutines in %v\n", count, time.Since(start))
}
