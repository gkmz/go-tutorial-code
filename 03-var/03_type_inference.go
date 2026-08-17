package main

import "fmt"

// 类型推导示例
func typeInference() {
	fmt.Println("=== 类型推导 ===")

	// 整数 → int
	a := 10
	fmt.Printf("a = %d, type: %T\n", a, a)

	// 浮点数 → float64
	b := 3.14
	fmt.Printf("b = %f, type: %T\n", b, b)

	// 字符串 → string
	c := "Go"
	fmt.Printf("c = %s, type: %T\n", c, c)

	// 布尔 → bool
	d := true
	fmt.Printf("d = %v, type: %T\n", d, d)

	// 多变量推导
	x, y := 10, 20
	fmt.Printf("x = %d, type: %T\n", x, x)
	fmt.Printf("y = %d, type: %T\n", y, y)

	// 函数返回值推导
	name := getName()
	fmt.Printf("name = %s, type: %T\n", name, name)

	// 切片推导
	nums := []int{1, 2, 3}
	fmt.Printf("nums = %v, type: %T\n", nums, nums)

	// map 推导
	scores := map[string]int{"Go": 100, "Java": 90}
	fmt.Printf("scores = %v, type: %T\n", scores, scores)
}

func getName() string {
	return "Go"
}

// 不能重新推导类型
func noReInference() {
	// 第一次声明
	var x int = 10
	fmt.Printf("x = %d, type: %T\n", x, x)

	// 重新赋值（可以改变值，但不能改变类型）
	x = 20
	fmt.Printf("x = %d, type: %T\n", x, x)

	// 用 := 重新声明会报错
	// x := "hello"  // 不支持 编译错误：no new variables
}

func main() {
	typeInference()

	fmt.Println("\n=== 不能重新推导 ===")
	noReInference()
}
