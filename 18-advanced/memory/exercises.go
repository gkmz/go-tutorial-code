package main

import "sync"

// growSlice 使用动态扩容构造切片，用于和预分配版本进行 benchmark 对比。
func growSlice(size int) []int {
	values := make([]int, 0)
	for i := 0; i < size; i++ {
		values = append(values, i)
	}
	return values
}

// preallocateSlice 预先分配底层数组，减少切片扩容和复制次数。
func preallocateSlice(size int) []int {
	values := make([]int, 0, size)
	for i := 0; i < size; i++ {
		values = append(values, i)
	}
	return values
}

// boundedCache 是练习题中的有界缓存参考实现。
// 它使用 FIFO 淘汰策略，重点演示缓存容量必须有明确上限。
type boundedCache struct {
	mu       sync.Mutex
	maxSize  int
	values   map[string][]byte
	keyOrder []string
}

// newBoundedCache 创建一个容量固定的缓存。
func newBoundedCache(maxSize int) *boundedCache {
	if maxSize < 1 {
		maxSize = 1
	}
	return &boundedCache{
		maxSize: maxSize,
		values:  make(map[string][]byte, maxSize),
	}
}

// Set 写入缓存，并在超过容量时淘汰最早写入的键。
func (c *boundedCache) Set(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.values[key]; !exists {
		c.keyOrder = append(c.keyOrder, key)
	}
	c.values[key] = append([]byte(nil), value...)

	for len(c.keyOrder) > c.maxSize {
		oldest := c.keyOrder[0]
		c.keyOrder = c.keyOrder[1:]
		delete(c.values, oldest)
	}
}

// Get 获取缓存值，并返回副本，避免调用方修改缓存内部数据。
func (c *boundedCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, exists := c.values[key]
	if !exists {
		return nil, false
	}
	return append([]byte(nil), value...), true
}
