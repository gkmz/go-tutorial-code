package main

import "fmt"

// 基础结构体
type Animal struct {
	Name string
	Age  int
}

// Animal 的方法
func (a Animal) Speak() string {
	return "Some sound"
}

// 嵌入 Animal
type Dog struct {
	Animal // 匿名嵌入
	Breed  string
}

// Dog 的方法
func (d Dog) Bark() string {
	return "Woof!"
}

// 重写 Speak 方法
func (d Dog) Speak() string {
	return "Woof! Woof!"
}

func main() {
	// 使用嵌入
	d := Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Golden Retriever",
	}

	// 直接访问嵌入字段（字段提升）
	fmt.Println("Name:", d.Name)
	fmt.Println("Age:", d.Age)
	fmt.Println("Breed:", d.Breed)

	// 调用嵌入的方法
	fmt.Println("Dog speaks:", d.Speak())           // Woof! Woof!（Dog 的方法）
	fmt.Println("Animal speaks:", d.Animal.Speak()) // Some sound（Animal 的方法）

	// 调用 Dog 自己的方法
	fmt.Println("Dog barks:", d.Bark())
}
