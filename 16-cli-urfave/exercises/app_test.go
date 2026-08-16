package exercises

import (
	"bytes"
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
