// Package exercises 提供错误处理章节练习的参考实现。
package exercises

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ErrNotFound 表示目标资源不存在。
var ErrNotFound = errors.New("not found")

// Divide 执行整数除法并返回除零错误。
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// RangeError 表示值超出允许范围。
type RangeError struct {
	Min   int
	Max   int
	Value int
}

// Error 返回范围错误文本。
func (e *RangeError) Error() string {
	return fmt.Sprintf("value %d is outside [%d, %d]", e.Value, e.Min, e.Max)
}

// ValidateRange 校验 value 是否位于闭区间内。
func ValidateRange(value, min, max int) error {
	if value < min || value > max {
		return &RangeError{Min: min, Max: max, Value: value}
	}
	return nil
}

// SafeCall 执行函数并把同一 goroutine 中的 panic 转为 error。
func SafeCall(fn func()) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("panic: %v", value)
		}
	}()
	fn()
	return nil
}

// Retry 在 context 取消、成功或达到次数上限时结束。
func Retry(ctx context.Context, attempts int, wait time.Duration, fn func() error) error {
	if attempts <= 0 {
		return errors.New("attempts must be positive")
	}
	var lastErr error
	for attempt := 0; attempt < attempts; attempt++ {
		if attempt > 0 {
			timer := time.NewTimer(wait)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
			}
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if lastErr = fn(); lastErr == nil {
			return nil
		}
		wait *= 2
	}
	return fmt.Errorf("after %d attempts: %w", attempts, lastErr)
}
