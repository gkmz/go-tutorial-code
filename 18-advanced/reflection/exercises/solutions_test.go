package exercises

import (
	"reflect"
	"testing"
)

func TestCopyMatchingFields(t *testing.T) {
	type source struct {
		Name string
		Age  int
	}
	type destination struct {
		Name string
		Age  int
	}
	var target destination
	if err := CopyMatchingFields(&target, source{Name: "Hank", Age: 18}); err != nil {
		t.Fatalf("CopyMatchingFields() error = %v", err)
	}
	if target != (destination{Name: "Hank", Age: 18}) {
		t.Fatalf("target = %#v", target)
	}
}

func TestMissingRequiredHandlesNestedValuesAndCycles(t *testing.T) {
	type profile struct {
		Email string `json:"email" validate:"required"`
	}
	type node struct {
		Name    string  `json:"name" validate:"required"`
		Profile profile `json:"profile"`
		Next    *node   `json:"next"`
	}
	root := &node{Name: "root"}
	root.Next = root

	got, err := MissingRequired(root)
	if err != nil {
		t.Fatalf("MissingRequired() error = %v", err)
	}
	want := []string{"profile.email"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("MissingRequired() = %v, want %v", got, want)
	}
}
