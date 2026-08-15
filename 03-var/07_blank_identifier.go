package main

import "fmt"

// 匿名变量示例
func blankIdentifier() {
	fmt.Println("=== 匿名变量 _ ===")

	// 接收所有返回值
	name, age, email := getUser()
	fmt.Printf("User: %s, %d, %s\n", name, age, email)

	// 只接收需要的值
	_, onlyAge, _ := getUser()
	fmt.Printf("Only age: %d\n", onlyAge)

	// 忽略函数返回值
	getUser() // 不使用返回值

	// 在循环中使用
	scores := []struct {
		name  string
		score int
	}{
		{"Alice", 90},
		{"Bob", 85},
		{"Charlie", 92},
	}

	for _, s := range scores {
		fmt.Printf("%s: %d\n", s.name, s.score)
	}
}

func getUser() (name string, age int, email string) {
	return "Alice", 30, "alice@example.com"
}

// _ 的注意事项
func blankNote() {
	fmt.Println("\n=== _ 的注意事项 ===")

	// _ 不能读取值
	_ = 42 // 可以赋值，但值被丢弃

	// 赋值给 _ 等于丢弃
	x, _ := 10, 20 // x=10，20 被丢弃
	fmt.Printf("x = %d\n", x)

	// 在 map 中使用
	m := map[string]int{"a": 1, "b": 2}
	for k, _ := range m {
		fmt.Println("key:", k)
	}

	// 在通道中使用
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for range ch {
		// 只关心通道是否关闭，不关心值
	}
}

func main() {
	blankIdentifier()
	blankNote()
}
