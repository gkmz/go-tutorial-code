package race

import (
	"sync"
	"testing"
)

// TestCounter 验证互斥锁保护共享状态，使用 -race 可进一步检测数据竞争。
func TestCounter(t *testing.T) {
	var mu sync.Mutex
	count := 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); mu.Lock(); count++; mu.Unlock() }()
	}
	wg.Wait()
	if count != 100 {
		t.Fatalf("count = %d, want 100", count)
	}
}
