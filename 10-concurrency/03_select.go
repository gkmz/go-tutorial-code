package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// main 演示 select 的多路等待、超时、非阻塞接收和取消。

func main() {
	fmt.Println("=== 示例 1: 基本 Select ===")
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(40 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	fmt.Println("\n=== 示例 2: 超时控制 ===")
	ch3 := make(chan string, 1)

	go func() {
		time.Sleep(40 * time.Millisecond)
		ch3 <- "result"
	}()

	select {
	case result := <-ch3:
		fmt.Println("Got:", result)
	case <-time.After(10 * time.Millisecond):
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch5 := make(chan int)
	ch6 := make(chan int)
	var producers sync.WaitGroup
	producers.Add(2)

	go func() {
		defer producers.Done()
		for i := 0; i < 3; i++ {
			select {
			case ch5 <- i:
			case <-ctx.Done():
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()

	go func() {
		defer producers.Done()
		for i := 0; i < 3; i++ {
			select {
			case ch6 <- i * 10:
			case <-ctx.Done():
				return
			}
			time.Sleep(7 * time.Millisecond)
		}
	}()

	go func() {
		producers.Wait()
		close(ch5)
		close(ch6)
	}()
	for ch5 != nil || ch6 != nil {
		select {
		case msg1, ok := <-ch5:
			if !ok {
				ch5 = nil
				continue
			}
			fmt.Println("From ch5:", msg1)
		case msg2, ok := <-ch6:
			if !ok {
				ch6 = nil
				continue
			}
			fmt.Println("From ch6:", msg2)
		case <-ctx.Done():
			return
		}
		if ch5 == nil && ch6 == nil {
			break
		}
	}
	cancel()
}
