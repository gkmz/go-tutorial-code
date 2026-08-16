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
	area := TotalArea([]Shape{Rectangle{Width: 2, Height: 3}, Circle{Radius: 1}})
	if math.Abs(area-(6+math.Pi)) > 1e-9 {
		t.Fatalf("area = %v", area)
	}
	if DescribeType(1) != "int" || DescribeType([]int{}) != "unknown" {
		t.Fatal("unexpected type description")
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
