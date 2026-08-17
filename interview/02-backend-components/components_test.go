package components

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunBatchCancelsAfterError(t *testing.T) {
	var started atomic.Int32
	failure := errors.New("failed")
	jobs := []Job{
		func(context.Context) error { started.Add(1); return failure },
		func(ctx context.Context) error { started.Add(1); <-ctx.Done(); return ctx.Err() },
		func(context.Context) error { started.Add(1); return nil },
	}
	if err := RunBatch(context.Background(), 1, jobs); !errors.Is(err, failure) {
		t.Fatalf("RunBatch() error = %v, want %v", err, failure)
	}
	if got := started.Load(); got != 1 {
		t.Fatalf("started %d jobs, want 1", got)
	}
}

func TestTTLCache(t *testing.T) {
	cache := NewTTLCache[string, int]()
	cache.Set("key", 1, time.Minute)
	if got, ok := cache.Get("key"); !ok || got != 1 {
		t.Fatalf("Get() = %d, %v; want 1, true", got, ok)
	}
	cache.Set("expired", 2, 0)
	if _, ok := cache.Get("expired"); ok {
		t.Fatal("expired item was returned")
	}
}

func TestPaginate(t *testing.T) {
	page, err := Paginate([]int{1, 2, 3, 4, 5}, 2, 2, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Items) != 2 || page.Items[0] != 3 || page.TotalPages != 3 {
		t.Fatalf("unexpected page: %+v", page)
	}
	if _, err := Paginate([]int{1}, 1, 11, 10); err == nil {
		t.Fatal("expected page size validation error")
	}
}

func TestDeduperConcurrent(t *testing.T) {
	deduper := NewDeduper()
	var wg sync.WaitGroup
	var first atomic.Int32
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if deduper.FirstSeen("same-event") {
				first.Add(1)
			}
		}()
	}
	wg.Wait()
	if first.Load() != 1 {
		t.Fatalf("first seen count = %d, want 1", first.Load())
	}
}
