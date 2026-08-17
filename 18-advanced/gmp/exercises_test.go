package main

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunLimitedHonorsConcurrencyLimit(t *testing.T) {
	const limit = 3
	var active atomic.Int32
	var maximum atomic.Int32

	err := runLimited(context.Background(), 30, limit, func(int) {
		current := active.Add(1)
		for {
			observed := maximum.Load()
			if current <= observed || maximum.CompareAndSwap(observed, current) {
				break
			}
		}
		time.Sleep(time.Millisecond)
		active.Add(-1)
	})
	if err != nil {
		t.Fatalf("runLimited() error = %v", err)
	}
	if got := maximum.Load(); got > limit {
		t.Fatalf("maximum concurrency = %d, want at most %d", got, limit)
	}
}

func TestRunLimitedStopsAfterCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var completed atomic.Int32

	err := runLimited(ctx, 100, 2, func(int) {
		if completed.Add(1) == 1 {
			cancel()
		}
	})
	if err != context.Canceled {
		t.Fatalf("runLimited() error = %v, want context.Canceled", err)
	}
	if got := completed.Load(); got >= 100 {
		t.Fatalf("completed tasks = %d, cancellation did not stop the workload", got)
	}
}

func TestRunLimitedRejectsInvalidArguments(t *testing.T) {
	tests := []struct {
		name string
		call func() error
	}{
		{name: "nil context", call: func() error { return runLimited(nil, 1, 1, func(int) {}) }},
		{name: "negative tasks", call: func() error { return runLimited(context.Background(), -1, 1, func(int) {}) }},
		{name: "zero limit", call: func() error { return runLimited(context.Background(), 1, 0, func(int) {}) }},
		{name: "nil work", call: func() error { return runLimited(context.Background(), 1, 1, nil) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); err == nil {
				t.Fatal("runLimited() should reject invalid arguments")
			}
		})
	}
}
