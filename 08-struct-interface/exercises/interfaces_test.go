package exercises

import (
	"encoding/json"
	"math"
	"strings"
	"testing"
)

func TestAnimalAndEmbedding(t *testing.T) {
	if FormatAnimal(Dog{}) != "dog eats; dog sleeps" {
		t.Fatal("unexpected dog behavior")
	}
	person := Person{Name: "Go", Address: Address{City: "Hangzhou"}}
	if person.City != "Hangzhou" {
		t.Fatal("embedded field was not promoted")
	}
}

func TestShapesAndTypeSwitch(t *testing.T) {
	area := TotalArea([]Shape{Rectangle{Width: 2, Height: 3}, Circle{Radius: 1}, Triangle{A: 3, B: 4, C: 5}})
	if math.Abs(area-(6+math.Pi+6)) > 1e-9 {
		t.Fatalf("area = %v", area)
	}
	if DescribeType(1) != "int" || DescribeType([]int{}) != "unknown" {
		t.Fatal("unexpected type description")
	}
	if SafeEqual([]int{1}, []int{1}) {
		t.Fatal("non-comparable slices must not compare equal")
	}
}

func TestNilAndTags(t *testing.T) {
	if WrongError() == nil || CorrectError() != nil {
		t.Fatal("unexpected interface nil behavior")
	}
	data, err := MarshalUser()
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "secret") || decoded["name"] != "Go" {
		t.Fatalf("unexpected JSON: %s", data)
	}
}
