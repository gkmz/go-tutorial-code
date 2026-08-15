package main

import (
	"fmt"
	"time"
)

// 基础 Goroutine 示例

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== 示例 1: 基础 Goroutine ===")
	go sayHello()
	fmt.Println("Hello from main")
	time.Sleep(time.Second)

	fmt.Println("\n=== 示例 2: 并发执行 ===")
	go printNumbers()
	go printLetters()
	time.Sleep(time.Second)
	fmt.Println("\nDone")

	fmt.Println("\n=== 示例 3: 匿名函数 Goroutine ===")
	go func() {
		fmt.Println("Anonymous goroutine")
	}()

	go func(msg string) {
		fmt.Println(msg)
	}("Hello from anonymous")

	time.Sleep(time.Second)
}
