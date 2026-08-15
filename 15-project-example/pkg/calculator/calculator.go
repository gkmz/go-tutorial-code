// calculator/calculator.go
package calculator

import (
	"fmt"
)

// Calculator 计算器结构
type Calculator struct {
	history []string
}

// New 创建计算器实例
func New() *Calculator {
	return &Calculator{
		history: []string{},
	}
}

// Add 加法
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.addHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

// Subtract 减法
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

// Multiply 乘法
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addHistory(fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
	return result
}

// Divide 除法
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	c.addHistory(fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
	return result, nil
}

// History 获取历史记录
func (c *Calculator) History() []string {
	return c.history
}

// addHistory 添加历史记录（私有方法）
func (c *Calculator) addHistory(record string) {
	c.history = append(c.history, record)
}
