package exercises

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestDivideAndSetAge(t *testing.T) {
	if result, err := Divide(10, 2); err != nil || result != 5 {
		t.Fatalf("Divide result = %d, %v", result, err)
	}
	if _, err := Divide(1, 0); err == nil {
		t.Fatal("expected division by zero")
	}
	user := User{Name: "Go"}
	user.SetAge(18)
	if user.Age != 18 {
		t.Fatalf("age = %d, want 18", user.Age)
	}
	if user.DisplayName() != "Go" {
		t.Fatalf("DisplayName() = %q, want Go", user.DisplayName())
	}
	var nilUser *User
	if nilUser.DisplayName() != "<nil user>" {
		t.Fatalf("nil DisplayName() = %q, want <nil user>", nilUser.DisplayName())
	}
}

func TestFilterAndCounter(t *testing.T) {
	got := Filter([]int{1, 2, 3, 4}, func(value int) bool { return value%2 == 0 })
	if !reflect.DeepEqual(got, []int{2, 4}) {
		t.Fatalf("Filter result = %v", got)
	}
	first, second := NewCounter(), NewCounter()
	if first() != 1 || first() != 2 || second() != 1 {
		t.Fatal("counters do not have independent state")
	}
}

func TestMarkFirstSharesExpandedSlice(t *testing.T) {
	values := []int{1, 2}
	MarkFirst(values...)
	if !reflect.DeepEqual(values, []int{-1, 2}) {
		t.Fatalf("values after MarkFirst() = %v, want [-1 2]", values)
	}
}

func TestMeasureAndRangeError(t *testing.T) {
	current := time.Unix(0, 0)
	var elapsed time.Duration
	stop := Measure(func() time.Time { return current }, func(value time.Duration) { elapsed = value })
	current = current.Add(2 * time.Second)
	stop()
	if elapsed != 2*time.Second {
		t.Fatalf("elapsed = %v, want 2s", elapsed)
	}
	var rangeErr *RangeError
	if !errors.As(ValidateRange(11, 0, 10), &rangeErr) || rangeErr.Value != 11 {
		t.Fatalf("unexpected range error: %v", rangeErr)
	}
}
