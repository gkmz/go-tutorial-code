// calculator/calculator.go
package calculator

import (
	"fmt"
	"sync"
)

// Calculator 提供基本算术运算，并保存当前实例的计算历史。
type Calculator struct {
	mu      sync.RWMutex
	history []string
}

// New 创建一个空的计算器实例。
func New() *Calculator {
	return &Calculator{
		history: []string{},
	}
}

// Add 计算两个浮点数的和，并记录操作历史。
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.addHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

// Subtract 计算两个浮点数的差，并记录操作历史。
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

// Multiply 计算两个浮点数的积，并记录操作历史。
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addHistory(fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
	return result
}

// Divide 计算两个浮点数的商，并拒绝除数为零。
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	c.addHistory(fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
	return result, nil
}

// History 返回当前历史记录的副本，调用方修改返回值不会影响计算器内部状态。
func (c *Calculator) History() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return append([]string(nil), c.history...)
}

// addHistory 在锁保护下追加一条历史记录。
func (c *Calculator) addHistory(record string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.history = append(c.history, record)
}
