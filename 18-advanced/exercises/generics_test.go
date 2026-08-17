package exercises

import (
	"reflect"
	"testing"
)

func TestGenericExercises(t *testing.T) {
	if got := Reverse([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Fatalf("Reverse() = %v", got)
	}

	var queue Queue[string]
	queue.Enqueue("first")
	queue.Enqueue("second")
	if got, ok := queue.Dequeue(); !ok || got != "first" {
		t.Fatalf("Dequeue() = %q, %v", got, ok)
	}

	if got := Reduce([]int{1, 2, 3}, 0, func(sum, value int) int { return sum + value }); got != 6 {
		t.Fatalf("Reduce() = %d", got)
	}

	var set Set[string]
	set.Add("go")
	set.Add("go")
	set.Remove("missing")
	if !set.Contains("go") {
		t.Fatal("Set does not contain inserted value")
	}

	groups := GroupBy([]int{1, 2, 3, 4}, func(value int) string {
		if value%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if !reflect.DeepEqual(groups["even"], []int{2, 4}) {
		t.Fatalf("GroupBy() = %v", groups)
	}
}
