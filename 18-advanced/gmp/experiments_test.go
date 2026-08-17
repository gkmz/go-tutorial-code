package main

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func TestRunGoroutineBurst(t *testing.T) {
	const count = 1_000
	baseline := runtime.NumGoroutine()

	result, err := runGoroutineBurst(count)
	if err != nil {
		t.Fatalf("runGoroutineBurst() error = %v", err)
	}
	if result.Count != count {
		t.Fatalf("Count = %d, want %d", result.Count, count)
	}
	if result.PeakGoroutines < baseline+count {
		t.Fatalf("PeakGoroutines = %d, want at least %d", result.PeakGoroutines, baseline+count)
	}
	if result.HeapAtPeak == 0 {
		t.Fatal("HeapAtPeak should contain a runtime measurement")
	}
}

func TestRunGoroutineBurstRejectsInvalidCount(t *testing.T) {
	if _, err := runGoroutineBurst(0); err == nil {
		t.Fatal("runGoroutineBurst(0) should return an error")
	}
}

func TestRunCancelableCPUStopsAfterCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		runCancelableCPU(ctx, 1_000)
		close(done)
	}()

	cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("CPU task did not observe cancellation")
	}
}

func TestRunPreemptionExperiment(t *testing.T) {
	result, err := runPreemptionExperiment(50 * time.Millisecond)
	if err != nil {
		t.Fatalf("runPreemptionExperiment() error = %v", err)
	}
	if result.ObserverTicks == 0 {
		t.Fatal("observer did not get a scheduling opportunity")
	}
}

func TestTraceWorkload(t *testing.T) {
	if _, err := runGoroutineBurst(200); err != nil {
		t.Fatalf("runGoroutineBurst() error = %v", err)
	}
	if _, err := runTimerExperiment(time.Millisecond); err != nil {
		t.Fatalf("runTimerExperiment() error = %v", err)
	}
}
