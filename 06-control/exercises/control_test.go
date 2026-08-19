package exercises

import (
	"reflect"
	"testing"
)

func TestGradeAndMultiples(t *testing.T) {
	if Grade(95) != "A" || Grade(59) != "F" {
		t.Fatal("unexpected grade")
	}
	if got := MultiplesOfThree(1, 10); !reflect.DeepEqual(got, []int{3, 6, 9}) {
		t.Fatalf("multiples = %v", got)
	}
}

func TestFindAndTypeSwitch(t *testing.T) {
	row, column, found := FindInMatrix([][]int{{1, 2}, {3, 4}}, 4)
	if !found || row != 1 || column != 1 {
		t.Fatalf("position = %d, %d, %v", row, column, found)
	}
	if DescribeType(1) != "int" || DescribeType("go") != "string" || DescribeType([]int{}) != "unknown" {
		t.Fatal("unexpected type description")
	}
}

func TestIncrementAges(t *testing.T) {
	ages := []int{18, 20}
	IncrementAges(ages)
	if !reflect.DeepEqual(ages, []int{19, 21}) {
		t.Fatalf("ages = %v", ages)
	}
}

func TestStableMapKeysAndLabelSearch(t *testing.T) {
	if got := StableMapKeys(map[string]int{"b": 2, "a": 1}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("keys = %v, want [a b]", got)
	}
	row, column, found := FindWithLabel([][]int{{1, 2}, {3, 4}}, 4)
	if !found || row != 1 || column != 1 {
		t.Fatalf("label search = %d, %d, %v", row, column, found)
	}
}

func TestIsTypedNil(t *testing.T) {
	var pointer *int
	if !IsTypedNil(pointer) {
		t.Fatal("typed nil pointer was not detected")
	}
	if IsTypedNil(1) {
		t.Fatal("integer must not be typed nil")
	}
}
