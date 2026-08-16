package exercises

import "testing"

func TestPower(t *testing.T) {
	if got := Power(2, 3); got != 8 {
		t.Fatalf("Power() = %v, want 8", got)
	}
}

func TestSquareRootRejectsNegative(t *testing.T) {
	if _, err := SquareRoot(-1); err == nil {
		t.Fatal("SquareRoot(-1) should return an error")
	}
}
