// Package exercises 提供函数章节练习的参考实现。
package exercises

import (
	"errors"
	"fmt"
	"time"
)

// Divide 使用命名返回值执行整数除法。
func Divide(a, b int) (result int, err error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// User 表示练习中的用户。
type User struct {
	Name string
	Age  int
}

// SetAge 修改用户年龄。
func (u *User) SetAge(age int) {
	u.Age = age
}

// DisplayName 返回用户名称；nil 接收者返回明确的占位文本。
func (u *User) DisplayName() string {
	if u == nil {
		return "<nil user>"
	}
	return u.Name
}

// Filter 返回满足 predicate 的整数。
func Filter(values []int, predicate func(int) bool) []int {
	result := make([]int, 0, len(values))
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// MarkFirst 修改可变参数切片的第一个元素，用于演示展开切片后的共享关系。
func MarkFirst(values ...int) {
	if len(values) > 0 {
		values[0] = -values[0]
	}
}

// NewCounter 返回一个从 1 开始递增的闭包计数器。
func NewCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Measure 返回一个记录 elapsed 的清理函数。
func Measure(now func() time.Time, report func(time.Duration)) func() {
	start := now()
	return func() {
		report(now().Sub(start))
	}
}

// RangeError 表示值超出允许范围。
type RangeError struct {
	Min   int
	Max   int
	Value int
}

// Error 返回范围错误的文本描述。
func (e *RangeError) Error() string {
	return fmt.Sprintf("value %d is outside range [%d, %d]", e.Value, e.Min, e.Max)
}

// ValidateRange 检查 value 是否位于闭区间 [min, max]。
func ValidateRange(value, min, max int) error {
	if value < min || value > max {
		return &RangeError{Min: min, Max: max, Value: value}
	}
	return nil
}
