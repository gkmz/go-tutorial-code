// Package exercises 提供 Context 章节的练习参考实现。
package exercises

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

// FetchAll 并发调用 fetch，并在任意调用失败或 Context 取消时结束。
func FetchAll(ctx context.Context, names []string, fetch func(context.Context, string) (string, error)) (map[string]string, error) {
	group, groupCtx := errgroup.WithContext(ctx)
	results := make(map[string]string, len(names))
	var mu sync.Mutex
	for _, name := range names {
		name := name
		group.Go(func() error {
			value, err := fetch(groupCtx, name)
			if err != nil {
				return err
			}
			mu.Lock()
			results[name] = value
			mu.Unlock()
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

// RunWorkers 使用固定数量的 worker 处理任务，并响应 Context 取消。
func RunWorkers(ctx context.Context, workers int, jobs []int, handle func(context.Context, int) error) error {
	if workers <= 0 {
		return &InvalidWorkerCountError{Workers: workers}
	}
	group, groupCtx := errgroup.WithContext(ctx)
	input := make(chan int)
	group.Go(func() error {
		defer close(input)
		for _, job := range jobs {
			select {
			case input <- job:
			case <-groupCtx.Done():
				return groupCtx.Err()
			}
		}
		return nil
	})
	for i := 0; i < workers; i++ {
		group.Go(func() error {
			for {
				select {
				case <-groupCtx.Done():
					return groupCtx.Err()
				case job, ok := <-input:
					if !ok {
						return nil
					}
					if err := handle(groupCtx, job); err != nil {
						return err
					}
				}
			}
		})
	}
	return group.Wait()
}

// InvalidWorkerCountError 表示 worker 数量不是正数。
type InvalidWorkerCountError struct {
	Workers int
}

// Error 返回错误描述。
func (e *InvalidWorkerCountError) Error() string {
	return "worker count must be greater than zero"
}
