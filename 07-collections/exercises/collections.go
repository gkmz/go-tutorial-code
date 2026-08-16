// Package exercises 提供集合章节练习的参考实现。
package exercises

import "sync"

// WordCount 统计文本中每个单词的出现次数。
func WordCount(words []string) map[string]int {
	result := make(map[string]int, len(words))
	for _, word := range words {
		result[word]++
	}
	return result
}

// Unique 返回保持首次出现顺序的去重切片。
func Unique(values []int) []int {
	seen := make(map[int]struct{}, len(values))
	result := make([]int, 0, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

// Reverse 原地反转整数切片。
func Reverse(values []int) {
	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		values[left], values[right] = values[right], values[left]
	}
}

// Cache 是一个由读写锁保护的并发安全缓存。
type Cache struct {
	mu     sync.RWMutex
	values map[string]string
}

// NewCache 创建一个空缓存。
func NewCache() *Cache {
	return &Cache{values: make(map[string]string)}
}

// Get 读取缓存中的值。
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.values[key]
	return value, ok
}

// Set 写入缓存中的值。
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

// Delete 删除缓存中的值。
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.values, key)
}
