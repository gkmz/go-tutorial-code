// Package exercises 提供计算器扩展练习的核心逻辑。
package exercises

import (
	"errors"
	"math"
)

var errNegativeRoot = errors.New("cannot calculate square root of a negative number")

// Power 返回 base 的 exponent 次方。
func Power(base, exponent float64) float64 { return math.Pow(base, exponent) }

// SquareRoot 返回 value 的平方根。
func SquareRoot(value float64) (float64, error) {
	if value < 0 {
		return 0, errNegativeRoot
	}
	return math.Sqrt(value), nil
}
