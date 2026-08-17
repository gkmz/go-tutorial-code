package lru

import (
	"sync"
	"testing"
)

func TestCacheEvictsLeastRecentlyUsed(t *testing.T) {
	c := New[string, int](2)
	c.Set("a", 1)
	c.Set("b", 2)
	_, _ = c.Get("a")
	c.Set("c", 3)
	if _, ok := c.Get("b"); ok {
		t.Fatal("expected b to be evicted")
	}
}

func TestCacheCapacityZero(t *testing.T) {
	c := New[string, int](0)
	c.Set("a", 1)
	if c.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", c.Len())
	}
}

func TestCacheConcurrent(t *testing.T) {
	c := New[int, int](32)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			c.Set(value, value)
			_, _ = c.Get(value)
		}(i)
	}
	wg.Wait()
	if c.Len() > 32 {
		t.Fatalf("Len() = %d, want <= 32", c.Len())
	}
}
