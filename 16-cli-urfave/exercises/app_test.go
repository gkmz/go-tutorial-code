package exercises

import (
	"bytes"
	"errors"
	"testing"
)

func TestHelloCommand(t *testing.T) {
	app := NewApp()
	var output bytes.Buffer
	app.Writer = &output
	if err := app.Run([]string{"greet", "hello", "--name", "Ada"}); err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if got, want := output.String(), "Hello, Ada!\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestHelloCommandUpper(t *testing.T) {
	app := NewApp()
	var output bytes.Buffer
	app.Writer = &output

	if err := app.Run([]string{"greet", "hello", "--name", "Ada", "--upper"}); err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if got, want := output.String(), "HELLO, ADA!\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestCountCommandRejectsNegativeNumber(t *testing.T) {
	app := NewApp()
	if err := app.Run([]string{"greet", "count", "--number", "-1"}); !errors.Is(err, ErrNegativeCount) {
		t.Fatalf("Run() error = %v, want ErrNegativeCount", err)
	}
}
