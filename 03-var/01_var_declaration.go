package main

import "fmt"

// var 声明示例
func varDeclaration() {
	// 声明单个变量
	var name string
	name = "Go"
	fmt.Println("name:", name)

	// 声明时直接赋值（类型可省略）
	var age = 18
	fmt.Printf("age: %d, type: %T\n", age, age)

	// 批量声明
	var (
		x     int
		y     float64
		z     bool
		s     string
		bytes byte
		r     rune
	)
	fmt.Printf("x: %d, y: %f, z: %v, s: %q, bytes: %d, r: %d\n", x, y, z, s, bytes, r)
}

// := 声明示例
func shortDeclaration() {
	// 必须初始化，Go 自动推导类型
	message := "Hello, Go!"
	fmt.Println("message:", message, "type:", fmt.Sprintf("%T", message))

	// 多变量声明
	a, b := 10, 20
	fmt.Printf("a: %d, b: %d\n", a, b)

	// 接收函数返回值
	result, err := someFunc()
	fmt.Printf("result: %d, err: %v\n", result, err)
}

func someFunc() (int, error) {
	return 42, nil
}

// 演示 var 和 := 的区别
func varVsShort() {
	// 用 var 声明
	var name string = "Go"
	fmt.Println("var name:", name)

	// 用 := 声明
	version := 1.21
	fmt.Println(":= version:", version)

	// 错误示例：已声明的变量不能再用 :=
	// name := "Python"  // ❌ 编译错误
	// 正确做法：赋值
	name = "Python"
	fmt.Println("after assignment:", name)
}

func main() {
	fmt.Println("=== var 声明 ===")
	varDeclaration()

	fmt.Println("\n=== := 声明 ===")
	shortDeclaration()

	fmt.Println("\n=== var vs := ===")
	varVsShort()
}
