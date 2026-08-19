// Package receivers 演示值接收者和指针接收者的方法集差异。
package receivers

// Counter 保存一个可以修改的计数值。
type Counter struct {
	value int
}

// Value 使用值接收者读取当前计数。
func (c Counter) Value() int {
	return c.value
}

// Add 使用指针接收者修改原始 Counter。
func (c *Counter) Add(delta int) {
	c.value += delta
}

// ValueReader 描述只读计数能力。
type ValueReader interface {
	Value() int
}

// MutableCounter 描述读取和修改计数的能力。
type MutableCounter interface {
	Value() int
	Add(int)
}
