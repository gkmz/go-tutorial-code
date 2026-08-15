package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup 示例

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Println("=== 示例 1: 基本 WaitGroup ===")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")

	fmt.Println("\n=== 示例 2: 避免闭包陷阱 ===")
	var wg2 sync.WaitGroup

	// 错误示例（注释掉）
	// for i := 0; i < 5; i++ {
	//     wg2.Add(1)
	//     go func() {
	//         defer wg2.Done()
	//         fmt.Println(i) // 可能都打印 5
	//     }()
	// }

	// 正确示例：传参
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func(n int) {
			defer wg2.Done()
			fmt.Println(n)
		}(i)
	}

	wg2.Wait()
	fmt.Println("Done")
}
