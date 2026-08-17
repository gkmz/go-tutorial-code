package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDescribe(t *testing.T) {
	fields, err := Describe(User{Name: "Hank", Age: 18, Email: "hank@example.com"})
	if err != nil {
		t.Fatalf("Describe() error = %v", err)
	}
	want := []FieldDescription{
		{Name: "Name", JSONName: "name", Value: "Hank"},
		{Name: "Age", JSONName: "age", Value: 18},
		{Name: "Email", JSONName: "email", Value: "hank@example.com"},
	}
	if !reflect.DeepEqual(fields, want) {
		t.Fatalf("Describe() = %#v, want %#v", fields, want)
	}
}

func TestDescribeRejectsInvalidInputs(t *testing.T) {
	var user *User
	for _, input := range []any{nil, user, 42} {
		if _, err := Describe(input); err == nil {
			t.Fatalf("Describe(%#v) error = nil", input)
		}
	}
}

func TestSetField(t *testing.T) {
	user := User{Name: "before"}
	if err := SetField(&user, "Name", "after"); err != nil {
		t.Fatalf("SetField() error = %v", err)
	}
	if user.Name != "after" {
		t.Fatalf("user.Name = %q", user.Name)
	}
}

func TestSetFieldRejectsInvalidOperations(t *testing.T) {
	user := User{}
	tests := []struct {
		name   string
		target any
		field  string
		value  any
	}{
		{name: "non pointer", target: user, field: "Name", value: "Hank"},
		{name: "missing field", target: &user, field: "Missing", value: "Hank"},
		{name: "wrong type", target: &user, field: "Age", value: "18"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := SetField(test.target, test.field, test.value); err == nil {
				t.Fatal("SetField() error = nil")
			}
		})
	}
}

func TestValidateRequired(t *testing.T) {
	valid := User{Name: "Hank", Email: "hank@example.com"}
	if err := ValidateRequired(valid); err != nil {
		t.Fatalf("ValidateRequired() error = %v", err)
	}

	invalid := User{Name: "Hank"}
	err := ValidateRequired(invalid)
	if err == nil || !strings.Contains(err.Error(), "Email") {
		t.Fatalf("ValidateRequired() error = %v", err)
	}
}

var benchmarkDescriptionSink []FieldDescription

func BenchmarkDescribeReflect(b *testing.B) {
	user := User{Name: "Hank", Age: 18, Email: "hank@example.com"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchmarkDescriptionSink, _ = Describe(user)
	}
}

func BenchmarkDescribeDirect(b *testing.B) {
	user := User{Name: "Hank", Age: 18, Email: "hank@example.com"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchmarkDescriptionSink = []FieldDescription{
			{Name: "Name", JSONName: "name", Value: user.Name},
			{Name: "Age", JSONName: "age", Value: user.Age},
			{Name: "Email", JSONName: "email", Value: user.Email},
		}
	}
}
