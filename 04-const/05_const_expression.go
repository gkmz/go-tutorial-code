package main

import "fmt"

// 常量表达式示例
func constExpression() {
	fmt.Println("=== 常量表达式 ===")

	// 算术表达式
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)
	fmt.Printf("KB=%d, MB=%d, GB=%d, TB=%d\n", KB, MB, GB, TB)

	// 时间常量
	const (
		SecondsPerMinute = 60
		SecondsPerHour   = 60 * SecondsPerMinute
		SecondsPerDay    = 24 * SecondsPerHour
	)
	fmt.Printf("Seconds per day: %d\n", SecondsPerDay)

	// 圆周率相关
	const (
		Pi  = 3.141592653589793
		Pi2 = Pi * Pi
		Pi3 = Pi2 * Pi
	)
	fmt.Printf("Pi=%.10f, Pi²=%.10f, Pi³=%.10f\n", Pi, Pi2, Pi3)
}

// 复数常量
func complexConst() {
	fmt.Println("\n=== 复数常量 ===")

	const (
		c1 = 1 + 2i
		c2 = 3 - 4i
		c3 = c1 + c2
		c4 = c1 * c2
	)
	fmt.Printf("c1 = %v\n", c1)
	fmt.Printf("c2 = %v\n", c2)
	fmt.Printf("c1 + c2 = %v\n", c3)
	fmt.Printf("c1 * c2 = %v\n", c4)

	// 复数的实部和虚部
	fmt.Printf("real(c1) = %f, imag(c1) = %f\n", real(c1), imag(c1))
}

// 字符串常量
func stringConst() {
	fmt.Println("\n=== 字符串常量 ===")

	const (
		Prefix = "Go"
		Sep    = "_"
		Suffix = "Lang"
		Name   = Prefix + Sep + Suffix
	)
	fmt.Printf("Name = %s\n", Name)

	// 原始字符串（反引号）
	const rawString = `Line 1
Line 2
Line 3`
	fmt.Printf("Raw string:\n%s\n", rawString)
}

// 布尔常量
func boolConst() {
	fmt.Println("\n=== 布尔常量 ===")

	const (
		True  = true
		False = false
	)
	fmt.Printf("True = %v, False = %v\n", True, False)

	// 组合布尔
	and := True && False
	or := True || False
	not := !True
	fmt.Printf("True && False = %v\n", and)
	fmt.Printf("True || False = %v\n", or)
	fmt.Printf("!True = %v\n", not)
}

// 数组/切片常量（编译期确定大小）
func arrayConst() {
	fmt.Println("\n=== 数组常量 ===")

	const (
		Count = 5
	)

	// 使用常量作为数组大小
	arr := [Count]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v, length: %d\n", arr, len(arr))

	// 切片不能直接用常量定义，但可以用常量计算
	const size = 10
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i * i
	}
	fmt.Printf("Slice: %v\n", slice)
}

// 枚举值表达式
func enumExpression() {
	fmt.Println("\n=== 枚举值表达式 ===")

	// 状态码枚举
	type HTTPStatus int

	const (
		StatusOK HTTPStatus = 200 + iota
		StatusCreated
		StatusAccepted
		StatusNoContent
		StatusBadRequest
		StatusUnauthorized
		StatusForbidden
		StatusNotFound
	)

	fmt.Printf("StatusOK = %d\n", StatusOK)
	fmt.Printf("StatusNotFound = %d\n", StatusNotFound)

	// 颜色枚举（RGB 值）
	type Color int

	const (
		ColorRed   Color = iota // 0
		ColorGreen              // 1
		ColorBlue               // 2
	)

	// 使用枚举值作为索引
	colorNames := [...]string{"Red", "Green", "Blue"}
	colorValues := [...]int{0xFF0000, 0x00FF00, 0x0000FF}

	for i := ColorRed; i <= ColorBlue; i++ {
		fmt.Printf("%s: name=%s, value=0x%06X\n", i, colorNames[i], colorValues[i])
	}
}

func main() {
	constExpression()
	complexConst()
	stringConst()
	boolConst()
	arrayConst()
	enumExpression()
}
