package exercises

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestFetchAll(t *testing.T) {
	results, err := FetchAll(context.Background(), []string{"a", "b"}, func(ctx context.Context, name string) (string, error) {
		return "value-" + name, nil
	})
	if err != nil {
		t.Fatalf("FetchAll() error = %v", err)
	}
	if results["a"] != "value-a" || results["b"] != "value-b" {
		t.Fatalf("results = %#v", results)
	}
}

func TestFetchAllCancelsRemainingCalls(t *testing.T) {
	wantErr := errors.New("fetch failed")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := FetchAll(ctx, []string{"fail", "wait"}, func(ctx context.Context, name string) (string, error) {
		if name == "fail" {
			return "", wantErr
		}
		<-ctx.Done()
		return "", ctx.Err()
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("FetchAll() error = %v, want %v", err, wantErr)
	}
}

func TestRunWorkers(t *testing.T) {
	var mu sync.Mutex
	var got []int
	err := RunWorkers(context.Background(), 2, []int{1, 2, 3}, func(_ context.Context, job int) error {
		mu.Lock()
		got = append(got, job*2)
		mu.Unlock()
		return nil
	})
	if err != nil {
		t.Fatalf("RunWorkers() error = %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("handled jobs = %d, want 3", len(got))
	}
}

func TestRunWorkersRejectsInvalidCount(t *testing.T) {
	if err := RunWorkers(context.Background(), 0, nil, nil); err == nil {
		t.Fatal("RunWorkers() error = nil")
	}
}
