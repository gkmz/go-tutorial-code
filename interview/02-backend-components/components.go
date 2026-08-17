// Package components 提供 Go 后端面试中的组件级参考实现。
package components

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Job 是一个必须响应 Context 取消的任务。
type Job func(context.Context) error

// RunBatch 使用固定数量的 Worker 执行任务；首个错误会取消剩余任务。
func RunBatch(ctx context.Context, workers int, jobs []Job) error {
	if workers <= 0 {
		workers = 1
	}
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	queue := make(chan Job)
	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
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
						case errCh <- err:
						default:
						}
						cancel()
						return
					}
				}
			}
		}()
	}

	for _, job := range jobs {
		select {
		case <-runCtx.Done():
			break
		case queue <- job:
		}
		if runCtx.Err() != nil {
			break
		}
	}
	close(queue)
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return runCtx.Err()
	}
}

// TTLCache 是一个进程内并发安全、惰性删除的 TTL 缓存。
type TTLCache[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]ttlItem[V]
}

type ttlItem[V any] struct {
	value     V
	expiresAt time.Time
}

// NewTTLCache 创建一个空的 TTL 缓存。
func NewTTLCache[K comparable, V any]() *TTLCache[K, V] {
	return &TTLCache[K, V]{items: make(map[K]ttlItem[V])}
}

// Set 写入值；ttl 小于等于零时，该值立即过期。
func (c *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = ttlItem[V]{value: value, expiresAt: time.Now().Add(ttl)}
}

// Get 返回未过期的值；过期项会在读取时删除。
func (c *TTLCache[K, V]) Get(key K) (V, bool) {
	now := time.Now()
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()
	if !ok || !now.Before(item.expiresAt) {
		if ok {
			c.mu.Lock()
			current, stillPresent := c.items[key]
			if stillPresent && !now.Before(current.expiresAt) {
				delete(c.items, key)
			}
			c.mu.Unlock()
		}
		var zero V
		return zero, false
	}
	return item.value, true
}

// Delete 删除指定键。
func (c *TTLCache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Page 表示一页查询结果。
type Page[T any] struct {
	Items      []T
	Page       int
	PageSize   int
	Total      int
	TotalPages int
}

// Paginate 对内存数据执行安全分页，页码从 1 开始。
func Paginate[T any](items []T, page, pageSize, maxPageSize int) (Page[T], error) {
	if page < 1 || pageSize < 1 || maxPageSize < 1 {
		return Page[T]{}, errors.New("page and page size must be positive")
	}
	if pageSize > maxPageSize {
		return Page[T]{}, errors.New("page size exceeds maximum")
	}
	start := (page - 1) * pageSize
	if start >= len(items) {
		return Page[T]{Items: []T{}, Page: page, PageSize: pageSize, Total: len(items), TotalPages: totalPages(len(items), pageSize)}, nil
	}
	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}
	result := append([]T(nil), items[start:end]...)
	return Page[T]{Items: result, Page: page, PageSize: pageSize, Total: len(items), TotalPages: totalPages(len(items), pageSize)}, nil
}

func totalPages(total, pageSize int) int {
	if total == 0 {
		return 0
	}
	return (total + pageSize - 1) / pageSize
}

// Deduper 是单进程内的事件去重器。
type Deduper struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

// NewDeduper 创建一个空的事件去重器。
func NewDeduper() *Deduper {
	return &Deduper{seen: make(map[string]struct{})}
}

// FirstSeen 标记事件；第一次出现返回 true，重复事件返回 false。
func (d *Deduper) FirstSeen(eventID string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.seen[eventID]; ok {
		return false
	}
	d.seen[eventID] = struct{}{}
	return true
}
