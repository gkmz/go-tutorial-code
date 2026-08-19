package greeter

import "testing"

func TestMessage(t *testing.T) {
	if got, want := Message("Ada"), "Hello, Ada!"; got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}
