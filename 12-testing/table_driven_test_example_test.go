package math

import (
	"testing"
)

// TestAddTableDriven 演示表格驱动测试
func TestAddTableDrivenExample(t *testing.T) {
	// 1. 定义测试用例结构
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"正数相加", 1, 2, 3},
		{"负数相加", -1, -2, -3},
		{"零值测试", 0, 0, 0},
		{"混合测试", -1, 5, 4},
	}

	// 2. 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 避坑引导：
// - 使用 t.Run 可以隔离不同子测试，其中一个失败不影响其他。
// - 结构体切片让增加测试用例变得非常简单。
