package runtimeexperiments

import (
	"context"
	"io"
	"testing"
)

func TestEscapeToHeap(t *testing.T) {
	value := EscapeToHeap(42)
	if value == nil || *value != 42 {
		t.Fatalf("value = %v, want pointer to 42", value)
	}
}

func TestAllocateBytes(t *testing.T) {
	if got := len(AllocateBytes(64)); got != 64 {
		t.Fatalf("len = %d, want 64", got)
	}
}

func TestCaptureCPUProfile(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := CaptureCPUProfile(ctx, io.Discard, func() {}); err != nil {
		t.Fatalf("CaptureCPUProfile() error = %v", err)
	}
}
