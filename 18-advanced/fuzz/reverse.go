// Package fuzz 提供 Go 原生模糊测试教程使用的边界处理示例。
package fuzz

import (
	"errors"
	"slices"
	"unicode/utf8"
)

// ErrInvalidUTF8 表示输入不是有效的 UTF-8 字符串。
var ErrInvalidUTF8 = errors.New("input is not valid UTF-8")

// ReverseBytes 返回输入字节的反转副本，不修改调用方提供的切片。
func ReverseBytes(input []byte) []byte {
	output := slices.Clone(input)
	for left, right := 0, len(output)-1; left < right; left, right = left+1, right-1 {
		output[left], output[right] = output[right], output[left]
	}
	return output
}

// ReverseUTF8 按 Unicode 码点反转有效 UTF-8 字符串。
func ReverseUTF8(input string) (string, error) {
	if !utf8.ValidString(input) {
		return "", ErrInvalidUTF8
	}
	runes := []rune(input)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes), nil
}
