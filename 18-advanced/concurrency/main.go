package main

import (
	"context"
	"fmt"
	"sync"
)

// workerPool 使用固定数量的 worker 处理任务，并通过 Context 支持取消。
func workerPool(ctx context.Context, workers int, jobs []int) []int {
	input := make(chan int)
	output := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case n, ok := <-input:
					if !ok {
						return
					}
					select {
					case output <- n * n:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	go func() {
		defer close(input)
		for _, n := range jobs {
			select {
			case input <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() { wg.Wait(); close(output) }()
	var result []int
	for n := range output {
		result = append(result, n)
	}
	return result
}

func main() { fmt.Println(workerPool(context.Background(), 2, []int{1, 2, 3, 4})) }
