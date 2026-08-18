package main

// Map 将 values 中的每个元素转换为 U，并返回新的切片。
func Map[T any, U any](values []T, convert func(T) U) []U {
	result := make([]U, len(values))
	for index, value := range values {
		result[index] = convert(value)
	}
	return result
}

// Filter 返回满足 predicate 的元素副本，不修改输入切片。
func Filter[T any](values []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(values))
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// Calculator 保存同一种数值类型的计算结果。
type Calculator[T Number] struct {
	history []T
}

// Add 计算两个数的和并记录结果。
func (c *Calculator[T]) Add(left, right T) T {
	result := left + right
	c.history = append(c.history, result)
	return result
}

// Subtract 计算左操作数减右操作数并记录结果。
func (c *Calculator[T]) Subtract(left, right T) T {
	result := left - right
	c.history = append(c.history, result)
	return result
}

// Multiply 计算两个数的乘积并记录结果。
func (c *Calculator[T]) Multiply(left, right T) T {
	result := left * right
	c.history = append(c.history, result)
	return result
}

// History 返回计算历史的副本，避免调用方修改内部切片。
func (c *Calculator[T]) History() []T {
	return append([]T(nil), c.history...)
}

// Average 返回历史结果的平均值。整数类型会执行整数除法。
func (c *Calculator[T]) Average() T {
	if len(c.history) == 0 {
		var zero T
		return zero
	}
	var total T
	for _, value := range c.history {
		total += value
	}
	return total / T(len(c.history))
}
