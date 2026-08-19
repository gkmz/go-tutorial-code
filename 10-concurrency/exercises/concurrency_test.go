package exercises

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"testing"
)

func TestChannelExercises(t *testing.T) {
	if got := SendAndReceive(42); got != 42 {
		t.Fatalf("SendAndReceive = %d, want 42", got)
	}
	if got := CloseAndDrain([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("CloseAndDrain = %v", got)
	}
}

func TestSafeCounter(t *testing.T) {
	var counter SafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	if counter.Value() != 100 {
		t.Fatalf("counter = %d, want 100", counter.Value())
	}
}

func TestRunPoolAndFanIn(t *testing.T) {
	results := RunPool(context.Background(), []int{1, 2, 3, 4}, 2)
	sort.Ints(results)
	if !reflect.DeepEqual(results, []int{2, 4, 6, 8}) {
		t.Fatalf("pool results = %v", results)
	}
	a, b := make(chan int, 1), make(chan int, 1)
	a <- 1
	b <- 2
	close(a)
	close(b)
	merged := []int{}
	for value := range FanIn(a, b) {
		merged = append(merged, value)
	}
	sort.Ints(merged)
	if !reflect.DeepEqual(merged, []int{1, 2}) {
		t.Fatalf("fan-in results = %v", merged)
	}
}

func TestRunPoolHonorsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if results := RunPool(ctx, []int{1, 2, 3}, 2); len(results) != 0 {
		t.Fatalf("cancelled pool returned %v, want no results", results)
	}
}
