// Package lru 提供一个并发安全的泛型 LRU 缓存。
package lru

import "sync"

type entry[K comparable, V any] struct {
	key        K
	value      V
	prev, next *entry[K, V]
}

// Cache 是一个使用 Map 和双向链表实现的并发安全 LRU 缓存。
type Cache[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	items    map[K]*entry[K, V]
	head     *entry[K, V]
	tail     *entry[K, V]
}

// New 创建一个指定容量的 LRU 缓存；负容量按零处理。
func New[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity < 0 {
		capacity = 0
	}
	return &Cache[K, V]{capacity: capacity, items: make(map[K]*entry[K, V], capacity)}
}

// Get 读取值并将命中节点移动到最近使用位置。
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	node, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.moveToFront(node)
	return node.value, true
}

// Set 写入值，超过容量时淘汰最久未使用的节点。
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.capacity == 0 {
		return
	}
	if node, ok := c.items[key]; ok {
		node.value = value
		c.moveToFront(node)
		return
	}
	node := &entry[K, V]{key: key, value: value}
	c.items[key] = node
	c.pushFront(node)
	if len(c.items) > c.capacity {
		old := c.tail
		c.remove(old)
		delete(c.items, old.key)
	}
}

// Len 返回当前缓存条目数。
func (c *Cache[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

func (c *Cache[K, V]) pushFront(node *entry[K, V]) {
	node.next = c.head
	if c.head != nil {
		c.head.prev = node
	} else {
		c.tail = node
	}
	c.head = node
}

func (c *Cache[K, V]) remove(node *entry[K, V]) {
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		c.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		c.tail = node.prev
	}
	node.prev, node.next = nil, nil
}

func (c *Cache[K, V]) moveToFront(node *entry[K, V]) {
	if c.head == node {
		return
	}
	c.remove(node)
	c.pushFront(node)
}
