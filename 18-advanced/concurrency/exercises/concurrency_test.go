package exercises

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSafeCacheConcurrentAccess(t *testing.T) {
	cache := NewSafeCache()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set("key", "value")
			_, _ = cache.Get("key")
		}(i)
	}
	wg.Wait()
	if value, ok := cache.Get("key"); !ok || value != "value" {
		t.Fatalf("cache value = %q, %v", value, ok)
	}
}

func TestOnceLoaderCachesError(t *testing.T) {
	var loader OnceLoader[string]
	var calls atomic.Int32
	initFn := func() (string, error) {
		calls.Add(1)
		return "", errors.New("load failed")
	}
	for i := 0; i < 3; i++ {
		if _, err := loader.Load(initFn); err == nil {
			t.Fatal("Load() error = nil")
		}
	}
	if calls.Load() != 1 {
		t.Fatalf("initializer calls = %d, want 1", calls.Load())
	}
}

func TestSessionStore(t *testing.T) {
	var store SessionStore
	store.Put("session-1", "user-1")
	value, ok := store.Load("session-1")
	if !ok || value != "user-1" {
		t.Fatalf("Load() = %v, %v", value, ok)
	}
	store.Delete("session-1")
	if _, ok := store.Load("session-1"); ok {
		t.Fatal("session still exists after Delete")
	}
}

func TestAtomicCounter(t *testing.T) {
	var counter AtomicCounter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	if got := counter.Value(); got != 100 {
		t.Fatalf("counter = %d, want 100", got)
	}
}

func TestAtomicConfigCopiesSnapshots(t *testing.T) {
	var config AtomicConfig
	config.Store(map[string]string{"mode": "safe"})
	loaded := config.Load()
	loaded["mode"] = "changed"
	if got := config.Load()["mode"]; got != "safe" {
		t.Fatalf("config mode = %q, want safe", got)
	}
}

func TestRunLimitedStopsOnError(t *testing.T) {
	wantErr := errors.New("job failed")
	jobs := []func(context.Context) error{
		func(context.Context) error { return wantErr },
		func(ctx context.Context) error {
			<-ctx.Done()
			return ctx.Err()
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := RunLimited(ctx, 2, jobs); !errors.Is(err, wantErr) {
		t.Fatalf("RunLimited() error = %v, want %v", err, wantErr)
	}
}

func TestRunLimitedRejectsInvalidLimit(t *testing.T) {
	if err := RunLimited(context.Background(), 0, nil); err == nil {
		t.Fatal("RunLimited() error = nil")
	}
}
