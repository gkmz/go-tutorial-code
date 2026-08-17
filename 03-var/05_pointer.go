package main

import "fmt"

// 指针基本示例
func pointerBasics() {
	fmt.Println("=== 指针基础 ===")

	x := 10
	p := &x // & 取地址，p 是 *int 类型

	fmt.Printf("x = %d\n", x)
	fmt.Printf("p = %p (address of x)\n", p)
	fmt.Printf("*p = %d (value pointed by p)\n", *p)

	// 通过指针修改值
	*p = 20
	fmt.Printf("After *p = 20: x = %d\n", x)
}

// 指针作为函数参数
func pointerAsParam() {
	fmt.Println("\n=== 指针作为函数参数 ===")

	a, b := 10, 20
	fmt.Printf("Before: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After: a=%d, b=%d\n", a, b)
}

// swap 交换两个整数的值
func swap(a, b *int) {
	// 交换指针指向的值
	*a, *b = *b, *a
}

// new 函数示例
func newExample() {
	fmt.Println("\n=== new 函数 ===")

	// new 创建指针并初始化为零值
	p := new(int) // *int，值为 0
	fmt.Printf("new(int): %d\n", *p)

	*p = 100
	fmt.Printf("After assignment: %d\n", *p)

	// 创建字符串指针
	strPtr := new(string)
	*strPtr = "Hello, Go!"
	fmt.Printf("new(string): %s\n", *strPtr)
}

// 指针数组和切片
func pointerArray() {
	fmt.Println("\n=== 指针数组和切片 ===")

	// 指针数组
	arr := [3]int{1, 2, 3}
	var ptrs [3]*int

	for i := 0; i < 3; i++ {
		ptrs[i] = &arr[i]
	}

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("ptrs: %v\n", ptrs)

	// 通过指针修改数组
	for i := 0; i < 3; i++ {
		*ptrs[i] *= 10
	}
	fmt.Printf("After modification: arr = %v\n", arr)

	// 指针切片
	slice := []int{1, 2, 3, 4, 5}
	var slicePtrs []*int

	for i := range slice {
		slicePtrs = append(slicePtrs, &slice[i])
	}

	fmt.Printf("slicePtrs: %v\n", slicePtrs)
}

// nil 指针
func nilPointer() {
	fmt.Println("\n=== nil 指针 ===")

	var p *int // 默认 nil
	fmt.Printf("nil pointer: %v\n", p)

	// nil 指针不能解引用
	// fmt.Println(*p)  // 不支持 panic: invalid memory address

	// 安全做法：检查 nil
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("pointer is nil, skip dereference")
	}

	// 返回 nil 指针的函数
	p = getNilPointer()
	if p == nil {
		fmt.Println("got nil pointer from function")
	}
}

func getNilPointer() *int {
	return nil
}

// 多级指针
func multiLevelPointer() {
	fmt.Println("\n=== 多级指针 ===")

	x := 10
	p := &x    // *int
	pp := &p   // **int
	ppp := &pp // ***int

	fmt.Printf("x = %d\n", x)
	fmt.Printf("*p = %d\n", *p)
	fmt.Printf("**pp = %d\n", **pp)
	fmt.Printf("***ppp = %d\n", ***ppp)

	// 通过多级指针修改值
	***ppp = 20
	fmt.Printf("After ***ppp = 20: x = %d\n", x)
}

func main() {
	pointerBasics()
	pointerAsParam()
	newExample()
	pointerArray()
	nilPointer()
	multiLevelPointer()
}
