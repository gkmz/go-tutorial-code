package main

import (
	"fmt"
	"sync"
)

// sayHello 演示最小的 goroutine 工作函数。

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%c ", i)
	}
}

// runWithWaitGroup 启动两个任务，并明确等待它们结束。
func runWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		printNumbers()
	}()
	go func() {
		defer wg.Done()
		printLetters()
	}()
	wg.Wait()
}

func main() {
	fmt.Println("=== 示例 1: 基础 Goroutine ===")
	var helloDone sync.WaitGroup
	helloDone.Add(1)
	go func() {
		defer helloDone.Done()
		sayHello()
	}()
	fmt.Println("Hello from main")
	helloDone.Wait()

	fmt.Println("\n=== 示例 2: 并发执行 ===")
	runWithWaitGroup()
	fmt.Println("\nDone")

	fmt.Println("\n=== 示例 3: 匿名函数 Goroutine ===")
	var anonymousDone sync.WaitGroup
	anonymousDone.Add(2)
	go func() {
		defer anonymousDone.Done()
		fmt.Println("Anonymous goroutine")
	}()

	go func(msg string) {
		defer anonymousDone.Done()
		fmt.Println(msg)
	}("Hello from anonymous")
	anonymousDone.Wait()
}
