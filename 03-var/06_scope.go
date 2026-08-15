package main

import "fmt"

// 全局变量
var globalVar = "全局变量"

func scopeExamples() {
	fmt.Println("=== 变量作用域 ===")

	// 局部变量
	localVar := "局部变量"
	fmt.Println(globalVar, localVar)

	// 块级作用域
	{
		innerVar := "内层变量"
		fmt.Println(globalVar, localVar, innerVar)
	}

	// 内层变量遮蔽外层
	shadowed := "外层"
	{
		shadowed := "内层"
		fmt.Println("inner scope:", shadowed)
	}
	fmt.Println("outer scope:", shadowed)
}

// 循环变量作用域（Go 1.22+）
func loopScope() {
	fmt.Println("\n=== 循环变量作用域 ===")

	// Go 1.22+：每次循环都是新变量
	for i := 0; i < 3; i++ {
		fmt.Printf("i = %d (address: %p)\n", i, &i)
	}

	// 切片元素是独立的
	slice := []int{1, 2, 3}
	for _, v := range slice {
		fmt.Printf("v = %d (address: %p)\n", v, &v)
	}
}

// 函数参数作用域
func paramScope(x int) {
	fmt.Println("\n=== 函数参数作用域 ===")
	fmt.Printf("x in paramScope: %d (address: %p)\n", x, &x)
}

func main() {
	scopeExamples()
	loopScope()

	// 函数参数
	y := 100
	fmt.Printf("y in main: %d (address: %p)\n", y, &y)
	paramScope(y)
}
