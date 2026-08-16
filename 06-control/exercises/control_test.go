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
