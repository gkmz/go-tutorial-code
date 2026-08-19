package main

import (
	"fmt"
)

// main 演示无缓冲、缓冲和关闭后的 channel 行为。

func main() {
	fmt.Println("=== 示例 1: 无缓冲 Channel ===")
	ch1 := make(chan string)
	done := make(chan struct{})

	go func() {
		msg := <-ch1
		fmt.Println("Received:", msg)
		close(done)
	}()

	ch1 <- "Hello"
	<-done

	fmt.Println("\n=== 示例 2: 缓冲 Channel ===")
	ch2 := make(chan int, 3)

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	fmt.Println("\n=== 示例 3: 关闭 Channel ===")
	ch3 := make(chan int, 3)

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	close(ch3)

	// 使用 ok 检查 channel 是否关闭
	for {
		value, ok := <-ch3
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received:", value)
	}

	fmt.Println("\n=== 示例 4: Range 遍历 Channel ===")
	ch4 := make(chan int, 3)
	ch4 <- 1
	ch4 <- 2
	ch4 <- 3
	close(ch4)

	for value := range ch4 {
		fmt.Println("Received:", value)
	}
}
