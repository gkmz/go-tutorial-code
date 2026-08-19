package receivers

import "testing"

func TestReceiverMethodSetsAndMutation(t *testing.T) {
	counter := Counter{}
	counter.Add(2)
	if got := counter.Value(); got != 2 {
		t.Fatalf("Value() = %d, want 2", got)
	}

	// Counter 和 *Counter 都满足只读接口，只有 *Counter 满足可变接口。
	var reader ValueReader = counter
	var mutable MutableCounter = &counter
	if reader.Value() != 2 || mutable.Value() != 2 {
		t.Fatal("unexpected interface values")
	}
}
