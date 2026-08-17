package algorithms

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if got := TwoSum([]int{2, 7, 11, 15}, 9); !reflect.DeepEqual(got, []int{0, 1}) {
		t.Fatalf("TwoSum() = %v, want [0 1]", got)
	}
	if got := TwoSum([]int{1, 2}, 8); got != nil {
		t.Fatalf("TwoSum() = %v, want nil", got)
	}
}

func TestMergeIntervals(t *testing.T) {
	input := []Interval{{1, 3}, {2, 6}, {8, 10}, {10, 12}}
	want := []Interval{{1, 6}, {8, 12}}
	if got := MergeIntervals(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("MergeIntervals() = %v, want %v", got, want)
	}
	if !reflect.DeepEqual(input, []Interval{{1, 3}, {2, 6}, {8, 10}, {10, 12}}) {
		t.Fatal("MergeIntervals modified input")
	}
}

func TestLongestUniqueSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"", 0},
		{"你好世界你", 4},
	}
	for _, test := range tests {
		if got := LongestUniqueSubstring(test.input); got != test.want {
			t.Errorf("LongestUniqueSubstring(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Value: 1, Left: &TreeNode{Value: 2}, Right: &TreeNode{Value: 3}}
	want := [][]int{{1}, {2, 3}}
	if got := LevelOrder(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("LevelOrder() = %v, want %v", got, want)
	}
	if got := LevelOrder(nil); got != nil {
		t.Fatalf("LevelOrder(nil) = %v, want nil", got)
	}
}
