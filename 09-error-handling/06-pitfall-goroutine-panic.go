package main

import (
	"fmt"
	"time"
)

// 坑 3：在 goroutine 中 panic 导致程序崩溃

// 不支持 错误：goroutine 中的 panic 无法被外部 recover
func unsafeGoroutine() {
	go func() {
		panic("goroutine panic") // 整个程序崩溃
	}()
}

// 已完成 正确：在 goroutine 内部处理 panic
func safeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in goroutine: %v\n", r)
			}
		}()
		fn()
	}()
}

// 更好的做法：返回错误而不是 panic
func safeGoWithError(fn func() error) <-chan error {
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()

		if err := fn(); err != nil {
			errChan <- err
		}
	}()
	return errChan
}

func main() {
	fmt.Println("=== 坑 3：goroutine 中的 panic ===")

	// 错误示例（会导致程序崩溃，所以注释掉）
	// fmt.Println("\n--- 错误做法（会崩溃）---")
	// unsafeGoroutine()
	// time.Sleep(time.Second)
	// fmt.Println("This won't print")

	// 正确示例 1：在 goroutine 内部 recover
	fmt.Println("\n--- 正确做法 1：内部 recover ---")
	safeGo(func() {
		panic("goroutine panic")
	})
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Program continues")

	// 正确示例 2：返回错误
	fmt.Println("\n--- 正确做法 2：返回错误 ---")
	errChan := safeGoWithError(func() error {
		panic("goroutine panic")
	})
	if err := <-errChan; err != nil {
		fmt.Println("Error from goroutine:", err)
	}
	fmt.Println("Program continues")

	// 正常情况
	fmt.Println("\n--- 正常情况 ---")
	errChan = safeGoWithError(func() error {
		fmt.Println("Goroutine working...")
		return nil
	})
	if err := <-errChan; err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Goroutine completed successfully")
	}

	// 教训：
	// 1. goroutine 中的 panic 无法被外部 recover
	// 2. 必须在 goroutine 内部处理 panic
	// 3. 更好的做法是返回错误而不是 panic
}
