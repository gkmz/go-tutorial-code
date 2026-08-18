package exercises

import (
	"sync"
	"testing"
)

func TestAccountCheckAndWithdrawIsAtomic(t *testing.T) {
	account := NewAccount(100)
	var successes int
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if account.Withdraw(80) {
				mu.Lock()
				successes++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if successes != 1 || account.Balance() != 20 {
		t.Fatalf("successes = %d, balance = %d", successes, account.Balance())
	}
}

func TestSignalAfterWrite(t *testing.T) {
	var value int
	done := SignalAfterWrite(&value, 42)
	<-done
	if value != 42 {
		t.Fatalf("value = %d, want 42", value)
	}
}
