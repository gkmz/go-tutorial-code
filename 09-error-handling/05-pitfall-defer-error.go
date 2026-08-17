package main

import (
	"fmt"
	"os"
)

// 坑 1：忘记检查 defer 中的错误

// 不支持 错误：忽略 Close 的错误
func processFileWrong(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // 忽略了 Close 的错误

	// 处理文件...
	fmt.Println("Processing file...")
	return nil
}

// 已完成 正确：检查 Close 的错误
func processFileRight(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil && err == nil {
			err = closeErr // 如果没有其他错误，返回 Close 错误
		}
	}()

	// 处理文件...
	fmt.Println("Processing file...")
	return nil
}

// 更好的做法：使用命名返回值
func processFileBest(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close file: %w", cerr)
		}
	}()

	// 处理文件...
	fmt.Println("Processing file...")
	return nil
}

func main() {
	fmt.Println("=== 坑 1：忘记检查 defer 中的错误 ===")

	// 创建一个测试文件
	testFile := "test.txt"
	f, _ := os.Create(testFile)
	f.WriteString("test content")
	f.Close()
	defer os.Remove(testFile)

	fmt.Println("\n--- 错误做法 ---")
	err := processFileWrong(testFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success (but Close error was ignored)")
	}

	fmt.Println("\n--- 正确做法 ---")
	err = processFileRight(testFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success (Close error was checked)")
	}

	fmt.Println("\n--- 最佳做法 ---")
	err = processFileBest(testFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success (with proper error wrapping)")
	}

	// 教训：
	// 1. defer 中的函数也可能返回错误
	// 2. 使用命名返回值可以在 defer 中修改返回值
	// 3. 如果已经有错误，Close 错误可以忽略（或记录日志）
}
