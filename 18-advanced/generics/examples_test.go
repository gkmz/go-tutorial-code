package main

import (
	"reflect"
	"testing"
)

func TestMapAndFilter(t *testing.T) {
	values := []int{1, 2, 3, 4}
	if got := Map(values, func(value int) string { return string(rune('0' + value)) }); !reflect.DeepEqual(got, []string{"1", "2", "3", "4"}) {
		t.Fatalf("Map() = %v", got)
	}
	if got := Filter(values, func(value int) bool { return value%2 == 0 }); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Fatalf("Filter() = %v", got)
	}
}

func TestCalculatorCopiesHistory(t *testing.T) {
	var calculator Calculator[int]
	calculator.Add(10, 20)
	calculator.Subtract(50, 20)
	calculator.Multiply(5, 6)
	history := calculator.History()
	history[0] = 99
	if got := calculator.History()[0]; got != 30 {
		t.Fatalf("History()[0] = %d, want 30", got)
	}
	if got := calculator.Average(); got != 30 {
		t.Fatalf("Average() = %d, want 30", got)
	}
}
