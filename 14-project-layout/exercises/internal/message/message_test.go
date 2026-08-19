package message

import "testing"

func TestFormat(t *testing.T) {
	for _, test := range []struct {
		input string
		want  string
	}{
		{input: "Ada", want: "Hello, Ada!"},
		{input: "", want: "Hello, Go!"},
	} {
		got := Format(test.input)
		if got != test.want {
			t.Errorf("Format(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestFarewell(t *testing.T) {
	for _, test := range []struct {
		input string
		want  string
	}{
		{input: "Ada", want: "Goodbye, Ada!"},
		{input: "", want: "Goodbye, Go!"},
	} {
		got := Farewell(test.input)
		if got != test.want {
			t.Errorf("Farewell(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
