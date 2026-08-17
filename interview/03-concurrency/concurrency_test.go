package concurrency

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestCounterConcurrent(t *testing.T) {
	var counter Counter
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Add(1)
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	if got := counter.Load(); got != 1000 {
		t.Fatalf("counter = %d, want 1000", got)
	}
}

func TestConsume(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	input := make(chan int, 3)
	input <- 1
	input <- 2
	input <- 3
	close(input)
	var total atomic.Int64
	if err := Consume(ctx, input, 2, func(value int) { total.Add(int64(value)) }); err != nil {
		t.Fatalf("Consume() error = %v", err)
	}
	if total.Load() != 6 {
		t.Fatalf("total = %d, want 6", total.Load())
	}
}

func TestGroupSharesResult(t *testing.T) {
	group := NewGroup[string, int]()
	var calls atomic.Int32
	fn := func(context.Context) (int, error) {
		calls.Add(1)
		time.Sleep(10 * time.Millisecond)
		return 42, nil
	}
	results := make(chan int, 2)
	for i := 0; i < 2; i++ {
		go func() {
			value, err := group.Do(context.Background(), "key", fn)
			if err == nil {
				results <- value
			}
		}()
	}
	if <-results != 42 || <-results != 42 {
		t.Fatal("unexpected shared result")
	}
	if calls.Load() != 1 {
		t.Fatalf("fn calls = %d, want 1", calls.Load())
	}
}

func TestGroupPropagatesError(t *testing.T) {
	group := NewGroup[string, int]()
	want := errors.New("upstream")
	_, err := group.Do(context.Background(), "key", func(context.Context) (int, error) { return 0, want })
	if !errors.Is(err, want) {
		t.Fatalf("error = %v, want %v", err, want)
	}
}

func TestRunnerStopsTasks(t *testing.T) {
	runner := NewRunner(context.Background())
	stopped := make(chan struct{})
	runner.Go(func(ctx context.Context) {
		<-ctx.Done()
		close(stopped)
	})
	runner.Stop()
	select {
	case <-stopped:
	case <-time.After(time.Second):
		t.Fatal("task did not stop")
	}
}
