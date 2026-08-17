package ratelimiter

import (
	"context"
	"testing"
	"time"
)

func TestLimiterUsesBurstTokens(t *testing.T) {
	limiter := New(100, 2)
	defer limiter.Close()
	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	if err := limiter.Wait(ctx); err == nil {
		t.Fatal("expected third wait to time out")
	}
}

func TestLimiterHonorsCancellation(t *testing.T) {
	limiter := New(1, 1)
	defer limiter.Close()
	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := limiter.Wait(ctx); err != context.Canceled {
		t.Fatalf("Wait() error = %v, want context.Canceled", err)
	}
}
