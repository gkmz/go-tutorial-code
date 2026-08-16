package exercises

import (
	"strings"
	"testing"
	"time"
)

func TestTextAndNumbers(t *testing.T) {
	if NormalizeText("  Go   stdlib ") != "Go stdlib" {
		t.Fatal("text was not normalized")
	}
	if value, err := ParseNumber("42"); err != nil || value != 42 {
		t.Fatalf("parsed value = %d, err = %v", value, err)
	}
	if _, err := ParseNumber("bad"); err == nil {
		t.Fatal("expected parse error")
	}
}

func TestJSONSortAndRandom(t *testing.T) {
	data, err := EncodeEvent(Event{ID: 1, Name: "Go", CreatedAt: time.Unix(0, 0), Secret: "hidden"})
	if err != nil || strings.Contains(string(data), "hidden") {
		t.Fatalf("unexpected JSON: %s, err = %v", data, err)
	}
	values := []int{3, 1, 2}
	if StableSearch(values, 2) != 1 {
		t.Fatalf("sorted values = %v", values)
	}
	if got, want := RandomSequence(42, 3), RandomSequence(42, 3); !equalInts(got, want) {
		t.Fatal("fixed seed should be reproducible")
	}
}

func equalInts(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
