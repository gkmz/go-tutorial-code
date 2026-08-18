package main

import (
	"errors"
	"testing"
)

type handler struct{}

func (handler) Handle(input string) (string, error) {
	if input == "fail" {
		return "", errors.New("requested failure")
	}
	return "handled:" + input, nil
}

func (handler) Wrong(int) string { return "wrong" }

func TestCallStringMethod(t *testing.T) {
	got, err := CallStringMethod(handler{}, "Handle", "request")
	if err != nil || got != "handled:request" {
		t.Fatalf("CallStringMethod() = %q, %v", got, err)
	}
	if _, err := CallStringMethod(handler{}, "Handle", "fail"); err == nil {
		t.Fatal("CallStringMethod() did not return the method error")
	}
}

func TestCallStringMethodRejectsInvalidSignatures(t *testing.T) {
	for _, name := range []string{"Missing", "Wrong"} {
		if _, err := CallStringMethod(handler{}, name, "request"); err == nil {
			t.Fatalf("CallStringMethod(%q) error = nil", name)
		}
	}
}
