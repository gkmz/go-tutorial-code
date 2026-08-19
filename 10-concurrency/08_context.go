package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker 演示在工作间隙监听取消信号，避免固定 Sleep 延迟退出。

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping: %v\n", id, ctx.Err())
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Printf("Worker %d working\n", id)
		}
	}
}

// main 演示超时取消和手动取消如何结束 goroutine。
func main() {
	fmt.Println("=== 示例 1: 使用 Context 超时控制 ===")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var workers sync.WaitGroup
	for i := 1; i <= 3; i++ {
		workers.Add(1)
		go func(id int) {
			defer workers.Done()
			worker(ctx, id)
		}(i)
	}

	<-ctx.Done()
	workers.Wait()
	fmt.Println("All workers stopped")

	fmt.Println("\n=== 示例 2: 手动取消 ===")
	ctx2, cancel2 := context.WithCancel(context.Background())

	var taskDone sync.WaitGroup
	taskDone.Add(1)
	go func() {
		defer taskDone.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("Task cancelled")
				return
			case <-time.After(100 * time.Millisecond):
				fmt.Println("Task running...")
			}
		}
	}()

	time.Sleep(250 * time.Millisecond)
	cancel2() // 手动取消
	taskDone.Wait()
}
