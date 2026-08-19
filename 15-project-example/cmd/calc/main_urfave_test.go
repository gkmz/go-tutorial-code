//go:build urfave

package main

import "testing"

func TestParseTwoNumbersRejectsInvalidInput(t *testing.T) {
	if _, _, err := parseTwoNumbers([]string{"10", "bad"}); err == nil {
		t.Fatal("parseTwoNumbers() error = nil, want error")
	}
}
