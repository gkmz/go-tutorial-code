package main

import "testing"

func TestBoundedCacheEvictsOldestValue(t *testing.T) {
	cache := newBoundedCache(2)
	cache.Set("first", []byte("1"))
	cache.Set("second", []byte("2"))
	cache.Set("third", []byte("3"))

	if _, ok := cache.Get("first"); ok {
		t.Fatal("the oldest cache entry should be evicted")
	}
	if value, ok := cache.Get("third"); !ok || string(value) != "3" {
		t.Fatalf("Get(third) = %q, %v", value, ok)
	}
}

func TestBoundedCacheReturnsValueCopy(t *testing.T) {
	cache := newBoundedCache(1)
	cache.Set("key", []byte("value"))

	value, ok := cache.Get("key")
	if !ok {
		t.Fatal("expected cached value")
	}
	value[0] = 'X'

	stored, ok := cache.Get("key")
	if !ok || string(stored) != "value" {
		t.Fatalf("cached value was modified through returned slice: %q", stored)
	}
}

func BenchmarkGrowingSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = growSlice(1024)
	}
}

func BenchmarkPreallocatedSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = preallocateSlice(1024)
	}
}

func BenchmarkBoundedCacheSet(b *testing.B) {
	cache := newBoundedCache(128)
	value := []byte("payload")
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Set("key", value)
	}
}
