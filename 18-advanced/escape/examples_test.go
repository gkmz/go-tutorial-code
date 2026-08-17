package main

import "testing"

var (
	benchmarkValueSink   User
	benchmarkPointerSink *User
)

func TestEscapeExamples(t *testing.T) {
	value := NewUserValue("value")
	if value.Name != "value" {
		t.Fatalf("NewUserValue() = %q", value.Name)
	}

	pointer := NewUserPointer("pointer")
	if pointer.Name != "pointer" {
		t.Fatalf("NewUserPointer() = %q", pointer.Name)
	}

	if numbers := MakeNumbers(3); len(numbers) != 3 {
		t.Fatalf("len(MakeNumbers()) = %d", len(numbers))
	}

	counter := NewCounter()
	if first, second := counter(), counter(); first != 1 || second != 2 {
		t.Fatalf("counter results = %d, %d", first, second)
	}
}

func BenchmarkReturnValue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchmarkValueSink = NewUserValue("Hank")
	}
}

func BenchmarkRetainedPointer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchmarkPointerSink = NewUserPointer("Hank")
	}
}
