package main

import (
	"errors"
	"fmt"
	"os"
)

// 错误包装示例

// 使用 %w 包装错误
func readConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// 使用 %w 包装原始错误
		return fmt.Errorf("failed to read config: %w", err)
	}
	defer file.Close()
	return nil
}

// 检查错误类型
func checkError() {
	err := readConfig("nonexistent.json")
	if err != nil {
		fmt.Println("Error:", err)

		// 使用 errors.Is 检查错误链中是否包含特定错误
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist (detected with errors.Is)")
		}

		// 使用 errors.As 提取特定类型的错误
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Printf("Failed path: %s (detected with errors.As)\n", pathErr.Path)
		}
	}
}

// 错误的做法：使用 %v 而不是 %w
func readConfigWrong(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// 不支持 使用 %v 会丢失原始错误
		return fmt.Errorf("failed to read config: %v", err)
	}
	defer file.Close()
	return nil
}

func checkErrorWrong() {
	err := readConfigWrong("nonexistent.json")
	if err != nil {
		fmt.Println("\nError:", err)

		// 无法检测到原始错误
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error: cannot detect os.ErrNotExist (because we used %v)")
		}
	}
}

func main() {
	fmt.Println("=== 正确的错误包装（使用 %w）===")
	checkError()

	fmt.Println("\n=== 错误的错误包装（使用 %v）===")
	checkErrorWrong()
}
