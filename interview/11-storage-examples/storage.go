// Package storageexamples 提供数据库、缓存和分布式锁面试的最小模型。
package storageexamples

import (
	"context"
	"errors"
	"sync"
)

// Cache 是 Cache Aside 示例所需的缓存接口。
type Cache[K comparable, V any] interface {
	Get(context.Context, K) (V, bool)
	Set(context.Context, K, V) error
}

// Loader 从主存储加载数据。
type Loader[K comparable, V any] func(context.Context, K) (V, error)

// CacheAside 先查缓存，未命中时读取主存储并回填缓存。
func CacheAside[K comparable, V any](ctx context.Context, cache Cache[K, V], key K, load Loader[K, V]) (V, error) {
	if value, ok := cache.Get(ctx, key); ok {
		return value, nil
	}
	value, err := load(ctx, key)
	if err != nil {
		var zero V
		return zero, err
	}
	if err := cache.Set(ctx, key, value); err != nil {
		// 缓存回填失败不应覆盖已经成功的主存储结果。
		return value, nil
	}
	return value, nil
}

// MemoryCache 是用于测试 Cache Aside 的内存缓存。
type MemoryCache[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]V
}

// NewMemoryCache 创建内存缓存。
func NewMemoryCache[K comparable, V any]() *MemoryCache[K, V] {
	return &MemoryCache[K, V]{items: make(map[K]V)}
}

// Get 读取缓存值。
func (c *MemoryCache[K, V]) Get(_ context.Context, key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.items[key]
	return value, ok
}

// Set 写入缓存值。
func (c *MemoryCache[K, V]) Set(_ context.Context, key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
	return nil
}

// FencingLock 是用于说明租约和 Fencing Token 的最小模型。
type FencingLock struct {
	mu      sync.Mutex
	owner   string
	token   uint64
	version uint64
}

// Acquire 获取新的递增 Token；调用方必须把 Token 写入下游数据。
func (l *FencingLock) Acquire(owner string) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.owner != "" {
		return 0, errors.New("lock is held")
	}
	l.version++
	l.owner, l.token = owner, l.version
	return l.token, nil
}

// Release 释放当前持有者的锁。
func (l *FencingLock) Release(owner string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.owner == owner {
		l.owner = ""
	}
}

// AcceptToken 判断写入方的 Token 是否不早于已接受版本。
func (l *FencingLock) AcceptToken(token uint64) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if token < l.version {
		return false
	}
	l.version = token
	return true
}
