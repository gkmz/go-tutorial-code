//go:build cgo

package main

import "testing"

func TestAdd(t *testing.T) {
	if got := add(2, 3); got != 5 {
		t.Fatalf("add(2, 3) = %d, want 5", got)
	}
}

func TestDuplicateMessage(t *testing.T) {
	got, err := duplicateMessage("hello")
	if err != nil {
		t.Fatalf("duplicateMessage() error = %v", err)
	}
	if got != "hello" {
		t.Fatalf("duplicateMessage() = %q, want hello", got)
	}
}

func TestDuplicateMessageWithStatus(t *testing.T) {
	got, err := duplicateMessageWithStatus("Go 1.25.1")
	if err != nil {
		t.Fatalf("duplicateMessageWithStatus() error = %v", err)
	}
	if got != "Go 1.25.1" {
		t.Fatalf("duplicateMessageWithStatus() = %q, want Go 1.25.1", got)
	}
}

func BenchmarkCAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(i, i)
	}
}

func BenchmarkGoAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i + i
	}
}
