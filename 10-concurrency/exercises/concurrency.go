// Package exercises 提供并发章节练习的参考实现。
package exercises

import (
	"context"
	"sync"
)

// SafeCounter 是一个由互斥锁保护的计数器。
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment 将计数器加一。
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value 返回计数器当前值。
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// RunPool 使用固定数量 worker 处理任务，并响应 context 取消。
func RunPool(ctx context.Context, jobs []int, workers int) []int {
	if workers <= 0 {
		return nil
	}
	input := make(chan int)
	output := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-input:
					if !ok {
						return
					}
					select {
					case output <- job * 2:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}
	go func() {
		defer close(input)
		for _, job := range jobs {
			select {
			case input <- job:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		wg.Wait()
		close(output)
	}()
	var results []int
	for result := range output {
		results = append(results, result)
	}
	return results
}

// FanIn 合并多个只读 channel，并在全部输入关闭后关闭输出。
func FanIn[T any](channels ...<-chan T) <-chan T {
	output := make(chan T)
	var wg sync.WaitGroup
	for _, channel := range channels {
		wg.Add(1)
		go func(input <-chan T) {
			defer wg.Done()
			for value := range input {
				output <- value
			}
		}(channel)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}
