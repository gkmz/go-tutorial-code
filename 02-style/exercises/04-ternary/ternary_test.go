package ternary

import "testing"

type user struct {
	Name string
}

func TestTernary(t *testing.T) {
	if got := Ternary(true, 1, 2); got != 1 {
		t.Fatalf("integer result = %d, want 1", got)
	}
	if got := Ternary(false, "true", "false"); got != "false" {
		t.Fatalf("string result = %q, want %q", got, "false")
	}
	want := user{Name: "go"}
	if got := Ternary(true, want, user{Name: "other"}); got != want {
		t.Fatalf("struct result = %#v, want %#v", got, want)
	}
}
