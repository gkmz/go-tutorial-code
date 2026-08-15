package fuzz

import (
	"slices"
	"testing"
)

// Reverse 返回输入字节的反转副本。
func Reverse(input []byte) []byte {
	output := slices.Clone(input)
	for left, right := 0, len(output)-1; left < right; left, right = left+1, right-1 {
		output[left], output[right] = output[right], output[left]
	}
	return output
}

// FuzzReverse 验证反转两次可以恢复原值。
func FuzzReverse(f *testing.F) {
	f.Add([]byte("hello"))
	f.Fuzz(func(t *testing.T, input []byte) {
		if got := Reverse(Reverse(input)); !slices.Equal(got, input) {
			t.Fatalf("round trip failed")
		}
	})
}
