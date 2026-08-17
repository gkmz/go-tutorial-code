package semantics

import (
	"errors"
	"reflect"
	"testing"
)

func TestSliceValueAndSharedArray(t *testing.T) {
	values := make([]int, 1, 2)
	values[0] = 1
	MutateElement(values, 0, 2)
	if values[0] != 2 {
		t.Fatalf("values[0] = %d, want 2", values[0])
	}
	values = AppendValue(values, 3)
	if !reflect.DeepEqual(values, []int{2, 3}) {
		t.Fatalf("values = %v, want [2 3]", values)
	}
}

func TestTypedNil(t *testing.T) {
	if TypedNilError() == nil {
		t.Fatal("typed nil error should not equal nil interface")
	}
}

func TestDeferEvaluation(t *testing.T) {
	argument, closure := DeferEvaluation()
	if argument != 1 || closure != 2 {
		t.Fatalf("got argument=%d closure=%d, want 1 and 2", argument, closure)
	}
}

func TestWrappedError(t *testing.T) {
	if !errors.Is(WrapNotFound("user"), ErrNotFound) {
		t.Fatal("wrapped error should match ErrNotFound")
	}
}

func TestMapValues(t *testing.T) {
	got := MapValues([]int{1, 2, 3}, func(value int) string { return string(rune('0' + value)) })
	if !reflect.DeepEqual(got, []string{"1", "2", "3"}) {
		t.Fatalf("got %v", got)
	}
}

func TestInspectPointer(t *testing.T) {
	value := 1
	typeName, kind, canSet := InspectPointer(&value)
	if typeName != "*int" || kind != reflect.Int || !canSet {
		t.Fatalf("got %s, %s, %v", typeName, kind, canSet)
	}
}
