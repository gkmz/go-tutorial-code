package race

import "sync"

// Counter 使用互斥锁保护计数值，适合需要复合读写语义的场景。
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment 将计数值加一。
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value 返回当前计数值。
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
