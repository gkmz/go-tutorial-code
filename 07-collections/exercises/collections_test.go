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
