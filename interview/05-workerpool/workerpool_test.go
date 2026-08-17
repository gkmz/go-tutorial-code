package workerpool

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
)

func TestPoolRunsAllJobs(t *testing.T) {
	var count atomic.Int32
	jobs := make([]Job, 5)
	for i := range jobs {
		jobs[i] = func(context.Context) error {
			count.Add(1)
			return nil
		}
	}
	if err := New(2).Run(context.Background(), jobs); err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if count.Load() != 5 {
		t.Fatalf("ran %d jobs, want 5", count.Load())
	}
}

func TestPoolPropagatesError(t *testing.T) {
	want := errors.New("job failed")
	if err := New(1).Run(context.Background(), []Job{func(context.Context) error { return want }}); !errors.Is(err, want) {
		t.Fatalf("Run() error = %v, want %v", err, want)
	}
}

func TestPoolHonorsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := New(2).Run(ctx, []Job{func(context.Context) error { t.Fatal("job should not start"); return nil }}); !errors.Is(err, context.Canceled) {
		t.Fatalf("Run() error = %v, want context.Canceled", err)
	}
}
