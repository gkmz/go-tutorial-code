// Package exercises 提供测试章节练习的参考实现。
package exercises

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Normalize 将输入中的连续 Unicode 空白规范化为单个空格。
func Normalize(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// UserRepository 定义问候逻辑读取用户名所需的最小能力。
type UserRepository interface {
	GetName(id int) (string, error)
}

// Greeting 使用仓库中的用户名构造问候语。
func Greeting(repository UserRepository, id int) (string, error) {
	name, err := repository.GetName(id)
	if err != nil {
		return "", fmt.Errorf("读取用户 %d: %w", id, err)
	}
	return "Hello, " + name, nil
}

// Expired 判断当前时间是否已经到达或超过截止时间。
func Expired(now func() time.Time, deadline time.Time) bool {
	return !now().Before(deadline)
}

// Counter 是支持并发累加的计数器。
type Counter struct {
	mu    sync.Mutex
	value int
}

// Add 原子地增加计数值。
func (counter *Counter) Add(delta int) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.value += delta
}

// Value 返回加锁保护的当前计数值。
func (counter *Counter) Value() int {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.value
}

// ConcatPlus 使用加法拼接字符串。
func ConcatPlus(values []string) string {
	result := ""
	for _, value := range values {
		result += value
	}
	return result
}

// ConcatBuilder 使用 strings.Builder 拼接字符串。
func ConcatBuilder(values []string) string {
	var builder strings.Builder
	for _, value := range values {
		builder.WriteString(value)
	}
	return builder.String()
}
