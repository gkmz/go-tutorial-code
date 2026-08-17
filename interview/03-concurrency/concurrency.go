// Package concurrency 提供并发面试题的可运行参考实现。
package concurrency

import (
	"context"
	"sync"
	"sync/atomic"
)

// Counter 是一个使用原子操作实现的并发安全计数器。
type Counter struct {
	value atomic.Int64
}

// Add 增加计数器并返回增加后的值。
func (c *Counter) Add(delta int64) int64 {
	return c.value.Add(delta)
}

// Load 读取当前计数。
func (c *Counter) Load() int64 {
	return c.value.Load()
}

// Consume 启动固定数量 Worker，消费输入任务并在 Context 取消时退出。
func Consume(ctx context.Context, input <-chan int, workers int, handle func(int)) error {
	if workers <= 0 {
		workers = 1
	}
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case value, ok := <-input:
					if !ok {
						return
					}
					handle(value)
				}
			}
		}()
	}
	wg.Wait()
	return ctx.Err()
}

// Group 合并同一时间窗口内的相同 Key 请求。
// 它只负责单进程内合并，不负责缓存结果。
type Group[K comparable, V any] struct {
	mu    sync.Mutex
	calls map[K]*call[V]
}

type call[V any] struct {
	done  chan struct{}
	value V
	err   error
}

// NewGroup 创建一个请求合并器。
func NewGroup[K comparable, V any]() *Group[K, V] {
	return &Group[K, V]{calls: make(map[K]*call[V])}
}

// Do 执行同一 Key 的一次回源调用，其余调用方共享结果。
func (g *Group[K, V]) Do(ctx context.Context, key K, fn func(context.Context) (V, error)) (V, error) {
	g.mu.Lock()
	if existing, ok := g.calls[key]; ok {
		g.mu.Unlock()
		select {
		case <-existing.done:
			return existing.value, existing.err
		case <-ctx.Done():
			var zero V
			return zero, ctx.Err()
		}
	}
	current := &call[V]{done: make(chan struct{})}
	g.calls[key] = current
	g.mu.Unlock()

	current.value, current.err = fn(ctx)
	close(current.done)
	g.mu.Lock()
	delete(g.calls, key)
	g.mu.Unlock()
	return current.value, current.err
}

// Runner 管理一组可取消的后台任务，并等待所有任务退出。
type Runner struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewRunner 创建一个后台任务管理器。
func NewRunner(parent context.Context) *Runner {
	ctx, cancel := context.WithCancel(parent)
	return &Runner{ctx: ctx, cancel: cancel}
}

// Go 启动一个后台任务。任务必须在 ctx.Done() 后主动返回。
func (r *Runner) Go(task func(context.Context)) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		task(r.ctx)
	}()
}

// Stop 取消所有任务并等待它们退出。
func (r *Runner) Stop() {
	r.cancel()
	r.wg.Wait()
}
