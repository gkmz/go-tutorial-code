// Package workerpool 提供支持取消和错误传播的 Worker Pool。
package workerpool

import (
	"context"
	"sync"
)

// Job 是一个必须响应 ctx 取消的任务函数。
type Job func(context.Context) error

// Pool 使用固定数量的 Worker 执行任务。
type Pool struct {
	workers int
}

// New 创建 Worker Pool；非正数按一个 Worker 处理。
func New(workers int) *Pool {
	if workers <= 0 {
		workers = 1
	}
	return &Pool{workers: workers}
}

// Run 执行任务；首个错误会取消派生 Context，并等待已启动任务退出。
func (p *Pool) Run(ctx context.Context, jobs []Job) error {
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	queue := make(chan Job)
	errs := make(chan error, 1)
	var wg sync.WaitGroup
	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		wg.Go(func() {
			defer wg.Done()
			for {
				select {
				case <-runCtx.Done():
					return
				case job, ok := <-queue:
					if !ok {
						return
					}
					if err := job(runCtx); err != nil {
						select {
						case errs <- err:
						default:
						}
						cancel()
						return
					}
				}
			}
		})
	}

submitJobs:
	for _, job := range jobs {
		select {
		case <-runCtx.Done():
			break submitJobs
		case queue <- job:
		}
		if runCtx.Err() != nil {
			break
		}
	}
	close(queue)
	wg.Wait()
	select {
	case err := <-errs:
		return err
	default:
		return runCtx.Err()
	}
}
