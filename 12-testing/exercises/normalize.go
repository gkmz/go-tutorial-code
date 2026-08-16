// Package exercises 提供测试章节练习的参考实现。
package exercises

import "strings"

// Normalize 将输入中的连续空白规范化为单个空格。
func Normalize(value string) string {
	return strings.Join(strings.Fields(value), " ")
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
