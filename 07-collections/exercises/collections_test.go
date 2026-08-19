package exercises

import (
	"reflect"
	"testing"
)

func TestCollectionOperations(t *testing.T) {
	if got := WordCount([]string{"go", "go", "slice"}); !reflect.DeepEqual(got, map[string]int{"go": 2, "slice": 1}) {
		t.Fatalf("WordCount = %v", got)
	}
	if got := Unique([]int{1, 2, 1, 3, 2}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("Unique = %v", got)
	}
	values := []int{1, 2, 3}
	Reverse(values)
	if !reflect.DeepEqual(values, []int{3, 2, 1}) {
		t.Fatalf("Reverse = %v", values)
	}
	if got := EvenValues([]int{1, 2, 3, 4}); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Fatalf("EvenValues = %v", got)
	}
	source := []int{1, 2, 3}
	clone := Clone(source)
	clone[0] = 99
	if source[0] != 1 {
		t.Fatalf("Clone shares source: %v", source)
	}
	appended := LimitedAppend(source[:2], 9)
	if !reflect.DeepEqual(appended, []int{1, 2, 9}) || !reflect.DeepEqual(source, []int{1, 2, 3}) {
		t.Fatalf("LimitedAppend changed source: appended=%v source=%v", appended, source)
	}
	if value, ok := LookupZero(map[string]int{"zero": 0}, "zero"); !ok || value != 0 {
		t.Fatal("existing zero value was not found")
	}
	if _, ok := LookupZero(map[string]int{"zero": 0}, "missing"); ok {
		t.Fatal("missing key was reported as present")
	}
}

func TestCache(t *testing.T) {
	cache := NewCache()
	cache.Set("name", "Go")
	if value, ok := cache.Get("name"); !ok || value != "Go" {
		t.Fatal("cache value not found")
	}
	cache.Delete("name")
	if _, ok := cache.Get("name"); ok {
		t.Fatal("cache value still exists")
	}
}
