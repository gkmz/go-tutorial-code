// Package ternary 提供本章泛型条件选择练习的参考实现。
package ternary

// Ternary 根据 condition 返回 trueValue 或 falseValue。
// 两个候选值在调用前都会完成求值；包含副作用或昂贵计算时应改用 if。
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
