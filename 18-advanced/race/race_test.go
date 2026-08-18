package race

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounterConcurrentAccess(t *testing.T) {
	const goroutines = 32
	const increments = 100

	var counter Counter
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for worker := 0; worker < goroutines; worker++ {
		go func() {
			defer wg.Done()
			for i := 0; i < increments; i++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()

	if got, want := counter.Value(), goroutines*increments; got != want {
		t.Fatalf("Counter.Value() = %d, want %d", got, want)
	}
}

func TestRegistryConcurrentAccess(t *testing.T) {
	registry := NewRegistry()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		name := fmt.Sprintf("service-%d", i)
		wg.Add(2)
		go func() {
			defer wg.Done()
			registry.Register(name, "127.0.0.1")
		}()
		go func() {
			defer wg.Done()
			registry.Lookup(name)
		}()
	}
	wg.Wait()
}

func TestWatchdog(t *testing.T) {
	var watchdog Watchdog
	now := time.Unix(100, 0)
	watchdog.KeepAlive(now)
	if watchdog.Expired(now.Add(-time.Second)) {
		t.Fatal("Watchdog.Expired() = true for a fresh heartbeat")
	}
	if !watchdog.Expired(now.Add(time.Second)) {
		t.Fatal("Watchdog.Expired() = false after the deadline")
	}
}
