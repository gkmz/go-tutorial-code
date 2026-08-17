// Package exercises 提供并发进阶章节的练习参考实现。
package exercises

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// SafeCache 是一个由读写锁保护的字符串缓存。
type SafeCache struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewSafeCache 创建一个空的并发安全缓存。
func NewSafeCache() *SafeCache {
	return &SafeCache{data: make(map[string]string)}
}

// Get 读取键对应的值；键不存在时返回 false。
func (c *SafeCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

// Set 写入键值对。
func (c *SafeCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
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

// RunLimited 使用信号量限制任务的最大并发数。
func RunLimited(ctx context.Context, limit int64, jobs []func(context.Context) error) error {
	if limit <= 0 {
		return &InvalidLimitError{Limit: limit}
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(int(limit))
	sem := semaphore.NewWeighted(limit)
	for _, job := range jobs {
		job := job
		group.Go(func() error {
			if err := sem.Acquire(groupCtx, 1); err != nil {
				return err
			}
			defer sem.Release(1)
			return job(groupCtx)
		})
	}
	return group.Wait()
}

// InvalidLimitError 表示并发上限不是正数。
type InvalidLimitError struct {
	Limit int64
}

// Error 返回错误描述。
func (e *InvalidLimitError) Error() string {
	return "concurrency limit must be greater than zero"
}
