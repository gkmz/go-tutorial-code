package exercises

import (
	"bytes"
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBufferPool(t *testing.T) {
	pool := NewBufferPool(64)
	buffer := pool.Get()
	buffer.WriteString("request data")
	if !pool.Put(buffer) {
		t.Fatal("BufferPool.Put() rejected a small buffer")
	}
	if reused := pool.Get(); reused.Len() != 0 {
		t.Fatalf("reused buffer length = %d", reused.Len())
	}

	large := bytes.NewBuffer(make([]byte, 0, 128))
	if pool.Put(large) {
		t.Fatal("BufferPool.Put() retained an oversized buffer")
	}
}

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
	reads, writes := cache.Stats()
	if reads != 101 || writes != 100 {
		t.Fatalf("cache stats = (%d, %d), want (101, 100)", reads, writes)
	}
}

func TestMutexCounter(t *testing.T) {
	var counter MutexCounter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			counter.Add(1)
		})
	}
	wg.Wait()
	if got := counter.Value(); got != 1000 {
		t.Fatalf("counter = %d, want 1000", got)
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
	store.Put("session-2", "user-2")
	store.Clear()
	if _, ok := store.Load("session-2"); ok {
		t.Fatal("session still exists after Clear")
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

func TestRunWithSemaphoreLimitsConcurrency(t *testing.T) {
	var current atomic.Int32
	var maximum atomic.Int32
	jobs := make([]func(context.Context) error, 20)
	for i := range jobs {
		jobs[i] = func(context.Context) error {
			running := current.Add(1)
			defer current.Add(-1)
			for {
				observed := maximum.Load()
				if running <= observed || maximum.CompareAndSwap(observed, running) {
					break
				}
			}
			time.Sleep(time.Millisecond)
			return nil
		}
	}

	if err := RunWithSemaphore(context.Background(), 5, jobs); err != nil {
		t.Fatalf("RunWithSemaphore() error = %v", err)
	}
	if got := maximum.Load(); got > 5 {
		t.Fatalf("maximum concurrency = %d, want at most 5", got)
	}
}

func TestDownloadAllCancelsOtherDownloads(t *testing.T) {
	wantErr := errors.New("download failed")
	started := make(chan struct{})
	fetch := func(ctx context.Context, url string) error {
		if url == "bad" {
			<-started
			return wantErr
		}
		close(started)
		<-ctx.Done()
		return ctx.Err()
	}

	err := DownloadAll(context.Background(), 2, []string{"slow", "bad"}, fetch)
	if !errors.Is(err, wantErr) {
		t.Fatalf("DownloadAll() error = %v, want %v", err, wantErr)
	}
}

func TestDownloadAllRejectsNilFetcher(t *testing.T) {
	if err := DownloadAll(context.Background(), 1, nil, nil); err == nil {
		t.Fatal("DownloadAll() error = nil")
	}
}
