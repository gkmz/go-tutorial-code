package main

import (
	"fmt"
	"time"
)

// Select 多路复用示例

func main() {
	fmt.Println("=== 示例 1: 基本 Select ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	fmt.Println("\n=== 示例 2: 超时控制 ===")
	ch3 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "result"
	}()

	select {
	case result := <-ch3:
		fmt.Println("Got:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}

	fmt.Println("\n=== 示例 3: 非阻塞操作 ===")
	ch4 := make(chan int)

	select {
	case msg := <-ch4:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available")
	}

	fmt.Println("\n=== 示例 4: 监听多个 Channel ===")
	ch5 := make(chan int)
	ch6 := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			ch5 <- i
			time.Sleep(500 * time.Millisecond)
		}
		quit <- true
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch6 <- i * 10
			time.Sleep(700 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-ch5:
			fmt.Println("From ch5:", msg1)
		case msg2 := <-ch6:
			fmt.Println("From ch6:", msg2)
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}
