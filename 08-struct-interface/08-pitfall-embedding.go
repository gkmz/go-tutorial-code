package main

import "fmt"

// 坑 3：嵌入字段的方法集

type Animal struct {
	Name string
}

func (a *Animal) SetName(name string) {
	a.Name = name
	fmt.Println("SetName called, new name:", a.Name)
}

// 错误：值嵌入
type DogWrong struct {
	Animal // 值嵌入
}

// 正确：指针嵌入
type DogRight struct {
	*Animal // 指针嵌入
}

func main() {
	fmt.Println("=== 嵌入字段的方法集陷阱 ===")

	// 错误示例
	fmt.Println("\n--- 值嵌入（错误）---")
	d1 := DogWrong{Animal: Animal{Name: "Buddy"}}
	fmt.Println("Before SetName:", d1.Name)
	d1.SetName("Max")                      // 编译通过，但...
	fmt.Println("After SetName:", d1.Name) // 还是 "Buddy"！

	// 原因：DogWrong 值嵌入了 Animal，但 SetName 是指针接收者
	// 实际调用的是 (&d1.Animal).SetName()，修改的是临时指针

	// 正确示例
	fmt.Println("\n--- 指针嵌入（正确）---")
	d2 := DogRight{Animal: &Animal{Name: "Buddy"}}
	fmt.Println("Before SetName:", d2.Name)
	d2.SetName("Max")
	fmt.Println("After SetName:", d2.Name) // "Max"

	// 教训：
	// - 如果嵌入类型有指针接收者的方法，使用指针嵌入
	// - 或者将方法改为值接收者（如果语义允许）
}
