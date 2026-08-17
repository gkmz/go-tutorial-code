package main

import (
	"errors"
	"fmt"
	"os"
)

// 完整示例：综合运用错误处理

// 自定义错误类型
type FileError struct {
	Path string
	Op   string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%s %s: %v", e.Op, e.Path, e.Err)
}

func (e *FileError) Unwrap() error {
	return e.Err
}

// 哨兵错误
var (
	ErrEmptyFile    = errors.New("file is empty")
	ErrInvalidInput = errors.New("invalid input")
)

// 读取文件
func readFile(path string) ([]byte, error) {
	if path == "" {
		return nil, ErrInvalidInput
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, &FileError{
			Path: path,
			Op:   "read",
			Err:  err,
		}
	}

	if len(data) == 0 {
		return nil, ErrEmptyFile
	}

	return data, nil
}

// 处理文件
func processFile(path string) error {
	data, err := readFile(path)
	if err != nil {
		// 检查特定错误
		if errors.Is(err, ErrEmptyFile) {
			return fmt.Errorf("cannot process empty file: %w", err)
		}

		if errors.Is(err, ErrInvalidInput) {
			return fmt.Errorf("invalid file path: %w", err)
		}

		// 检查错误类型
		var fileErr *FileError
		if errors.As(err, &fileErr) {
			return fmt.Errorf("failed to access %s: %w", fileErr.Path, err)
		}

		return err
	}

	fmt.Printf("Successfully read %d bytes from %s\n", len(data), path)
	return nil
}

// 安全执行
func safeExecute(name string, fn func() error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %s: recovered from panic: %v\n", name, r)
		}
	}()

	fmt.Printf("\n=== %s ===\n", name)
	if err := fn(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	fmt.Println("=== 完整错误处理示例 ===")

	// 创建测试文件
	testFile := "test.txt"
	f, _ := os.Create(testFile)
	f.WriteString("test content")
	f.Close()
	defer os.Remove(testFile)

	// 创建空文件
	emptyFile := "empty.txt"
	os.Create(emptyFile)
	defer os.Remove(emptyFile)

	// 示例 1：正常文件
	safeExecute("Normal File", func() error {
		return processFile(testFile)
	})

	// 示例 2：空文件
	safeExecute("Empty File", func() error {
		return processFile(emptyFile)
	})

	// 示例 3：不存在的文件
	safeExecute("Non-existent File", func() error {
		return processFile("nonexistent.txt")
	})

	// 示例 4：无效输入
	safeExecute("Invalid Input", func() error {
		return processFile("")
	})

	// 示例 5：捕获 panic
	safeExecute("Panic Example", func() error {
		panic("unexpected error")
	})

	fmt.Println("\nProgram completed successfully")
}
