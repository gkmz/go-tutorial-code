package main

import (
	"fmt"
)

// getUserName 模拟获取用户名
func getUserName() (string, error) {
	return "VipUser", nil
}

func main() {
	fmt.Println("--- 坑 1：影子变量 (Shadowing) ---")
	shadowing()

	fmt.Println("\n--- 坑 2：类型系统的“洁癖” ---")
	typeSystem()

	fmt.Println("\n--- 坑 3：空指针的“自作多情” ---")
	nilPointer()
}

func shadowing() {
	var user = "Guest"
	isVip := true

	if isVip {
		// 不支持 坑在这里！用 := 声明了一个新的局部 user
		user, _ := getUserName()
		fmt.Println("内部用户:", user) // 打印 VipUser
	}
	fmt.Println("外部用户:", user) // 打印 Guest！外面的变量没变
}

func typeSystem() {
	var a int = 10
	var b int64 = 20

	// c := a + b // 不支持 编译报错：invalid operation (mismatched types int and int64)

	// 正确做法：显式强转
	c := int64(a) + b
	fmt.Printf("a(%T) + b(%T) = c(%T): %d\n", a, b, c, c)
}

func nilPointer() {
	var p *int
	// *p = 100 // 不支持 运行时 Panic: invalid memory address or nil pointer dereference

	if p == nil {
		fmt.Println("p 是 nil，直接解引用会 Panic！")
	}

	// 正确做法：初始化
	p = new(int)
	*p = 100
	fmt.Println("初始化后的 p:", *p)
}
