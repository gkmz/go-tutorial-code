package main

import "fmt"

// Number 定义支持加法的预声明数值类型及其自定义底层类型。
type Number interface{ ~int | ~int64 | ~float64 }

// Sum 返回切片中所有元素的和。
func Sum[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}), Sum([]float64{1.5, 2.5}))

	var calculator Calculator[int]
	calculator.Add(10, 20)
	fmt.Println(calculator.History(), calculator.Average())
}
