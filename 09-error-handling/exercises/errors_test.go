package exercises

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestErrors(t *testing.T) {
	if _, err := Divide(1, 0); err == nil {
		t.Fatal("expected divide error")
	}
	wrapped := errors.New("wrapped")
	if !errors.Is(errors.Join(ErrNotFound, wrapped), ErrNotFound) {
		t.Fatal("errors.Join lost sentinel")
	}
	var rangeErr *RangeError
	if !errors.As(ValidateRange(11, 0, 10), &rangeErr) || rangeErr.Value != 11 {
		t.Fatal("failed to extract RangeError")
	}
}

func TestSafeCallAndRetry(t *testing.T) {
	if err := SafeCall(func() { panic("boom") }); err == nil {
		t.Fatal("expected panic error")
	}
	attempts := 0
	err := Retry(context.Background(), 3, time.Millisecond, func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary")
		}
		return nil
	})
	if err != nil || attempts != 3 {
		t.Fatalf("retry result = %v, attempts = %d", err, attempts)
	}
}
