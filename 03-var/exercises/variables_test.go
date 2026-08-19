package exercises

import (
	"errors"
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	if a, b := Swap(10, 20); a != 20 || b != 10 {
		t.Fatalf("Swap result = %d, %d, want 20, 10", a, b)
	}

	a, b := 10, 20
	SwapInPlace(&a, &b)
	if a != 20 || b != 10 {
		t.Fatalf("SwapInPlace result = %d, %d, want 20, 10", a, b)
	}
}

func TestWeekday(t *testing.T) {
	if Monday != 0 || Sunday != 6 {
		t.Fatalf("weekday values = %d, %d, want 0, 6", Monday, Sunday)
	}
}

func TestBytesAndRunes(t *testing.T) {
	bytes, runes := BytesAndRunes("Go 中")
	if len(bytes) != 6 || len(runes) != 4 {
		t.Fatalf("lengths = %d bytes, %d runes, want 6, 4", len(bytes), len(runes))
	}
	if !reflect.DeepEqual(runes, []rune{'G', 'o', ' ', '中'}) {
		t.Fatalf("runes = %#v", runes)
	}
}

func TestIntegerBounds(t *testing.T) {
	int8Max, uint8Max, int32Max, int64Max := IntegerBounds()
	if int8Max != 127 || uint8Max != 255 || int32Max != 2147483647 || int64Max != 9223372036854775807 {
		t.Fatalf("unexpected integer bounds: %d, %d, %d, %d", int8Max, uint8Max, int32Max, int64Max)
	}
}

func TestReplaceUserUpdatesOuterValueAndPreservesError(t *testing.T) {
	got, err := ReplaceUser("guest", true, func() (string, error) {
		return "vip", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "vip" {
		t.Fatalf("ReplaceUser() = %q, want vip", got)
	}

	errLoad := errors.New("user unavailable")
	_, err = ReplaceUser("guest", true, func() (string, error) {
		return "", errLoad
	})
	if !errors.Is(err, errLoad) {
		t.Fatalf("ReplaceUser() error = %v, want wrapped load error", err)
	}
}
