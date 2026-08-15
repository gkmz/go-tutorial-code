package main

import (
	"fmt"
	"sync"
)

// Mutex 并发安全示例

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// 不安全的计数器（用于对比）
type UnsafeCounter struct {
	value int
}

func (c *UnsafeCounter) Increment() {
	c.value++
}

func main() {
	fmt.Println("=== 示例 1: 不安全的计数器 ===")
	var wg1 sync.WaitGroup
	unsafeCounter := &UnsafeCounter{}

	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			unsafeCounter.Increment()
		}()
	}

	wg1.Wait()
	fmt.Printf("Unsafe counter (可能不是 1000): %d\n", unsafeCounter.value)

	fmt.Println("\n=== 示例 2: 安全的计数器 ===")
	var wg2 sync.WaitGroup
	safeCounter := &Counter{}

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			safeCounter.Increment()
		}()
	}

	wg2.Wait()
	fmt.Printf("Safe counter (应该是 1000): %d\n", safeCounter.Value())
}
