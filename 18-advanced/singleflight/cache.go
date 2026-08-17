package main

import (
	"sync"
	"time"
)

type cacheEntry[V any] struct {
	value     V
	expiresAt time.Time
}

type ttlCache[K comparable, V any] struct {
	mu      sync.Mutex
	entries map[K]cacheEntry[V]
	now     func() time.Time
}

func newTTLCache[K comparable, V any]() *ttlCache[K, V] {
	return &ttlCache[K, V]{
		entries: make(map[K]cacheEntry[V]),
		now:     time.Now,
	}
}

func (c *ttlCache[K, V]) get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		var zero V
		return zero, false
	}
	if !c.now().Before(entry.expiresAt) {
		delete(c.entries, key)
		var zero V
		return zero, false
	}
	return entry.value, true
}

func (c *ttlCache[K, V]) set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	c.entries[key] = cacheEntry[V]{value: value, expiresAt: c.now().Add(ttl)}
	c.mu.Unlock()
}
