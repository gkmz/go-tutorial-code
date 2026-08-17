package main

import "fmt"

// 常量组
const (
	StatusOK = iota
	StatusPending
	StatusFailed
)

const (
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
	GB                    // 1073741824
)

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	// ============ 变量声明 ============
	fmt.Println("=== 变量声明 ===")

	var name string = "Go"
	version := 1.21
	fmt.Printf("name: %s, version: %.2f\n", name, version)

	// ============ 指针 ============
	fmt.Println("\n=== 指针 ===")

	ptr := &version
	fmt.Printf("version: %d, pointer: %p\n", version, ptr)
	*ptr = 1.22
	fmt.Printf("after modification: version = %.2f\n", version)

	// ============ 常量 ============
	fmt.Println("\n=== 常量 ===")

	fmt.Printf("Status: %d (OK), %d (Pending), %d (Failed)\n",
		StatusOK, StatusPending, StatusFailed)
	fmt.Printf("KB: %d, MB: %d, GB: %d\n", KB, MB, GB)

	// ============ 多重赋值和交换 ============
	fmt.Println("\n=== 多重赋值 ===")

	x, y := 10, 20
	fmt.Printf("Before: x=%d, y=%d\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("After swap: x=%d, y=%d\n", x, y)

	// 交换也可以直接写
	a, b := 100, 200
	a, b = b, a
	fmt.Printf("Direct swap: a=%d, b=%d\n", a, b)

	// ============ 匿名变量 ============
	fmt.Println("\n=== 匿名变量 ===")

	_, age := "Bob", 25
	fmt.Printf("Age: %d\n", age)

	// ============ 类型推导 ============
	fmt.Println("\n=== 类型推导 ===")

	num := 42
	fmt.Printf("num: %d, type: %T\n", num, num)

	str := "Hello"
	fmt.Printf("str: %s, type: %T\n", str, str)

	flag := true
	fmt.Printf("flag: %v, type: %T\n", flag, flag)

	// ============ 作用域 ============
	fmt.Println("\n=== 作用域 ===")

	outer := "outer"
	{
		inner := "inner"
		fmt.Println(outer, inner)
	}
	// fmt.Println(inner) // 不支持 编译错误
	fmt.Println(outer)

	fmt.Println("\n=== 完整示例完成 ===")
}
