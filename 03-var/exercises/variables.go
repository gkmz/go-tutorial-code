// Package exercises 提供变量章节练习的参考实现。
package exercises

import "fmt"

// Weekday 表示一周中的一天。
type Weekday int

const (
	// Monday 表示星期一。
	Monday Weekday = iota
	// Tuesday 表示星期二。
	Tuesday
	// Wednesday 表示星期三。
	Wednesday
	// Thursday 表示星期四。
	Thursday
	// Friday 表示星期五。
	Friday
	// Saturday 表示星期六。
	Saturday
	// Sunday 表示星期日。
	Sunday
)

// Swap 使用多重赋值交换两个整数的值。
func Swap(a, b int) (int, int) {
	return b, a
}

// SwapInPlace 使用指针直接交换调用方变量的值。
func SwapInPlace(a, b *int) {
	*a, *b = *b, *a
}

// BytesAndRunes 返回字符串的 UTF-8 字节切片和 rune 切片。
func BytesAndRunes(value string) ([]byte, []rune) {
	return []byte(value), []rune(value)
}

// IntegerBounds 返回常见整数类型的最大值。
func IntegerBounds() (int8, uint8, int32, int64) {
	return 1<<7 - 1, 1<<8 - 1, 1<<31 - 1, 1<<63 - 1
}

// ReplaceUser 在需要时更新外层用户值，并保留加载错误。
func ReplaceUser(current string, shouldReplace bool, load func() (string, error)) (string, error) {
	if !shouldReplace {
		return current, nil
	}

	// 使用赋值而不是短声明，确保更新的是当前函数作用域中的变量。
	var err error
	current, err = load()
	if err != nil {
		return "", fmt.Errorf("load replacement user: %w", err)
	}
	return current, nil
}
