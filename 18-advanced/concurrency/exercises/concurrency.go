// Package exercises 提供并发进阶章节的练习参考实现。
package exercises

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// SafeCache 是一个由读写锁保护的字符串缓存。
type SafeCache struct {
	mu     sync.RWMutex
	data   map[string]string
	reads  atomic.Uint64
	writes atomic.Uint64
}

// NewSafeCache 创建一个空的并发安全缓存。
func NewSafeCache() *SafeCache {
	return &SafeCache{data: make(map[string]string)}
}

// Get 读取键对应的值；键不存在时返回 false。
func (c *SafeCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.reads.Add(1)
	value, ok := c.data[key]
	return value, ok
}

// Set 写入键值对。
func (c *SafeCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	c.writes.Add(1)
}

// Stats 返回缓存从创建以来的读取和写入次数。
func (c *SafeCache) Stats() (reads, writes uint64) {
	return c.reads.Load(), c.writes.Load()
}

// OnceLoader 用 sync.Once 保证初始化函数最多执行一次。
type OnceLoader[T any] struct {
	once  sync.Once
	value T
	err   error
}

// Load 执行初始化函数并返回缓存的结果。
// 初始化函数返回错误后，后续调用仍会得到同一个错误，不会自动重试。
func (l *OnceLoader[T]) Load(initFn func() (T, error)) (T, error) {
	l.once.Do(func() {
		l.value, l.err = initFn()
	})
	return l.value, l.err
}

// SessionStore 使用 sync.Map 保存并发访问的会话数据。
type SessionStore struct {
	data sync.Map
}

// MutexCounter 是一个由互斥锁保护的整数计数器。
type MutexCounter struct {
	mu    sync.Mutex
	value int64
}

// Add 将计数器增加 delta。
func (c *MutexCounter) Add(delta int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

// Value 返回计数器当前值。
func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// AtomicCounter 是使用原子整数实现的并发安全计数器。
type AtomicCounter struct {
	value atomic.Int64
}

// Add 将计数器增加 delta。
func (c *AtomicCounter) Add(delta int64) {
	c.value.Add(delta)
}

// Value 返回计数器当前值。
func (c *AtomicCounter) Value() int64 {
	return c.value.Load()
}

// AtomicConfig 保存可以整体替换的配置快照。
type AtomicConfig struct {
	value atomic.Value
}

// Store 发布新的配置快照。所有快照必须使用相同的动态类型。
func (c *AtomicConfig) Store(config map[string]string) {
	copyOfConfig := make(map[string]string, len(config))
	for key, value := range config {
		copyOfConfig[key] = value
	}
	c.value.Store(copyOfConfig)
}

// Load 返回配置快照的副本，调用方可以安全修改返回值。
func (c *AtomicConfig) Load() map[string]string {
	value := c.value.Load().(map[string]string)
	copyOfConfig := make(map[string]string, len(value))
	for key, item := range value {
		copyOfConfig[key] = item
	}
	return copyOfConfig
}

// Put 保存会话值。
func (s *SessionStore) Put(id string, value any) {
	s.data.Store(id, value)
}

// Load 读取会话值。
func (s *SessionStore) Load(id string) (any, bool) {
	return s.data.Load(id)
}

// Delete 删除会话值。
func (s *SessionStore) Delete(id string) {
	s.data.Delete(id)
}

// Clear 删除全部会话。
func (s *SessionStore) Clear() {
	s.data.Clear()
}

// RunLimited 使用 errgroup.SetLimit 限制任务的最大并发数。
func RunLimited(ctx context.Context, limit int64, jobs []func(context.Context) error) error {
	if limit <= 0 {
		return &InvalidLimitError{Limit: limit}
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(int(limit))
	for _, job := range jobs {
		job := job
		group.Go(func() error {
			return job(groupCtx)
		})
	}
	return group.Wait()
}

// RunWithSemaphore 使用带权信号量限制同时占用资源的任务数量。
func RunWithSemaphore(ctx context.Context, limit int64, jobs []func(context.Context) error) error {
	if limit <= 0 {
		return &InvalidLimitError{Limit: limit}
	}

	group, groupCtx := errgroup.WithContext(ctx)
	sem := semaphore.NewWeighted(limit)
	for _, job := range jobs {
		// 在启动 Goroutine 前获取许可，避免创建大量等待中的 Goroutine。
		if err := sem.Acquire(groupCtx, 1); err != nil {
			if groupErr := group.Wait(); groupErr != nil {
				return groupErr
			}
			return err
		}

		job := job
		group.Go(func() error {
			defer sem.Release(1)
			return job(groupCtx)
		})
	}
	return group.Wait()
}

// DownloadAll 并发执行下载函数，并在任一下载失败后取消其他任务。
func DownloadAll(
	ctx context.Context,
	limit int64,
	urls []string,
	fetch func(context.Context, string) error,
) error {
	if fetch == nil {
		return &InvalidFetcherError{}
	}

	jobs := make([]func(context.Context) error, 0, len(urls))
	for _, url := range urls {
		url := url
		jobs = append(jobs, func(jobCtx context.Context) error {
			return fetch(jobCtx, url)
		})
	}
	return RunLimited(ctx, limit, jobs)
}

// InvalidLimitError 表示并发上限不是正数。
type InvalidLimitError struct {
	Limit int64
}

// Error 返回错误描述。
func (e *InvalidLimitError) Error() string {
	return "concurrency limit must be greater than zero"
}

// InvalidFetcherError 表示下载函数为空。
type InvalidFetcherError struct{}

// Error 返回错误描述。
func (e *InvalidFetcherError) Error() string {
	return "fetch function must not be nil"
}
