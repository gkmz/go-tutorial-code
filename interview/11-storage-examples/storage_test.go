package storageexamples

import (
	"context"
	"errors"
	"testing"
)

type testCache struct {
	value    string
	found    bool
	setError error
}

func (c *testCache) Get(context.Context, string) (string, bool) { return c.value, c.found }
func (c *testCache) Set(context.Context, string, string) error {
	c.value, c.found = "cached", true
	return c.setError
}

func TestCacheAside(t *testing.T) {
	cache := &testCache{}
	loaded := 0
	value, err := CacheAside(context.Background(), cache, "key", func(context.Context, string) (string, error) {
		loaded++
		return "from-db", nil
	})
	if err != nil || value != "from-db" || loaded != 1 {
		t.Fatalf("got value=%q err=%v loaded=%d", value, err, loaded)
	}
	cache.found = true
	if _, err := CacheAside(context.Background(), cache, "key", func(context.Context, string) (string, error) {
		t.Fatal("loader should not run on cache hit")
		return "", nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestCacheAsideLoadError(t *testing.T) {
	want := errors.New("database unavailable")
	_, err := CacheAside(context.Background(), &testCache{}, "key", func(context.Context, string) (string, error) { return "", want })
	if !errors.Is(err, want) {
		t.Fatalf("err = %v, want %v", err, want)
	}
}

func TestFencingLock(t *testing.T) {
	lock := &FencingLock{}
	first, err := lock.Acquire("a")
	if err != nil || first != 1 {
		t.Fatalf("first token = %d, err = %v", first, err)
	}
	if _, err := lock.Acquire("b"); err == nil {
		t.Fatal("second owner should be rejected")
	}
	lock.Release("a")
	second, err := lock.Acquire("b")
	if err != nil || second != 2 {
		t.Fatalf("second token = %d, err = %v", second, err)
	}
	if lock.AcceptToken(1) {
		t.Fatal("stale token should be rejected")
	}
}
