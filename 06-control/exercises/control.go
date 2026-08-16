// Package exercises 提供控制结构章节练习的参考实现。
package exercises

// Grade 根据分数返回等级。
func Grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// MultiplesOfThree 返回闭区间 [start, end] 中 3 的倍数。
func MultiplesOfThree(start, end int) []int {
	result := make([]int, 0)
	for value := start; value <= end; value++ {
		if value%3 != 0 {
			continue
		}
		result = append(result, value)
	}
	return result
}

// FindInMatrix 返回目标值第一次出现的行列位置。
func FindInMatrix(matrix [][]int, target int) (row, column int, found bool) {
	for i, values := range matrix {
		for j, value := range values {
			if value == target {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

// DescribeType 返回基础类型 switch 的描述。
func DescribeType(value any) string {
	switch typed := value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		return "float64"
	default:
		_ = typed
		return "unknown"
	}
}

// IncrementAges 使用索引修改切片中的结构体元素。
func IncrementAges(ages []int) {
	for i := range ages {
		ages[i]++
	}
}
