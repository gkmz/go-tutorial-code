package shadow

import (
	"errors"
	"testing"
)

func TestLoadCombinedPreservesExtraError(t *testing.T) {
	errExtra := errors.New("extra unavailable")
	_, err := LoadCombined(
		func() (int, error) { return 10, nil },
		func() (int, error) { return 0, errExtra },
	)
	if !errors.Is(err, errExtra) {
		t.Fatalf("LoadCombined() error = %v, want wrapped extra error", err)
	}
}

func TestLoadCombinedReturnsSum(t *testing.T) {
	got, err := LoadCombined(
		func() (int, error) { return 10, nil },
		func() (int, error) { return 5, nil },
	)
	if err != nil {
		t.Fatal(err)
	}
	if got != 15 {
		t.Fatalf("LoadCombined() = %d, want 15", got)
	}
}
