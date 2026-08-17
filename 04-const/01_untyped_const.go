package main

import "fmt"

// MyInt 演示定义类型后的常量类型检查。
type MyInt int

const (
	Small MyInt = iota
	Medium
	Large
)

// 无类型常量 vs 有类型常量
func untypedVsTyped() {
	fmt.Println("=== 无类型常量 vs 有类型常量 ===")

	// 有类型常量
	const piFloat32 float32 = 3.1415926
	const rInt int = 4
	// fmt.Println(piFloat32 * rInt)  // 不支持 编译错误：类型不匹配
	fmt.Printf("piFloat32 * float32(rInt) = %f\n", piFloat32*float32(rInt))

	// 无类型常量（推荐）
	const pi = 3.1415926 // 不指定类型
	const r = 4
	const area = pi * r * r // 已完成 直接运算，Go 自动处理类型
	fmt.Printf("area = %f (type: %T)\n", area, area)

	// 验证类型推导
	fmt.Printf("r type: %T, pi type: %T, area type: %T\n", r, pi, area)
}

// 无类型常量的灵活性
func untypedFlexibility() {
	fmt.Println("\n=== 无类型常量的灵活性 ===")

	const x = 10 // 无类型常量

	// 可以赋值给任何整数类型
	var i1 int = x
	var i2 int8 = x
	var i2_16 int16 = x
	var i3 int32 = x
	var i4 int64 = x
	var u1 uint = x
	var u2 uint8 = x

	fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n", i1, i2, i2_16, i3, i4)
	fmt.Printf("uint: %d, uint8: %d\n", u1, u2)

	// 可以赋值给浮点数
	var f1 float32 = x
	var f2 float64 = x
	fmt.Printf("float32: %f, float64: %f\n", f1, f2)

	// 可以赋值给复数
	var c complex128 = x
	fmt.Printf("complex128: %f\n", real(c))
}

// 有类型常量的必要性
func typedConstNecessity() {
	fmt.Println("\n=== 有类型常量的必要性 ===")

	// 函数参数类型安全
	checkSize(Small)
	checkSize(Medium)
	checkSize(Large)
	// checkSize(100)  // 不支持 编译错误：类型不匹配

	// 验证类型
	var x MyInt = Small
	fmt.Printf("Small type: %T, value: %d\n", x, x)
}

func checkSize(s MyInt) {
	if s == Small {
		fmt.Println("Small size")
	} else if s == Medium {
		fmt.Println("Medium size")
	} else if s == Large {
		fmt.Println("Large size")
	}
}

func main() {
	untypedVsTyped()
	untypedFlexibility()
	typedConstNecessity()
}
