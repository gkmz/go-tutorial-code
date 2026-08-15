package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// readSmallFile 读取小文件（几MB以内）
// 适合一次性读取整个文件内容
func readSmallFile(filename string) (string, error) {
	// ReadFile 一次性读取整个文件，返回 []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	// 将字节切片转换为字符串返回
	return string(data), nil
}

// readLargeFileLineByLine 逐行读取大文件
// 使用 bufio.Scanner，内存占用小，适合处理日志等大文件
func readLargeFileLineByLine(filename string) error {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	// defer 确保函数结束时关闭文件，避免资源泄漏
	defer file.Close()

	// 创建 Scanner，默认按行分割，每行最大 64KB
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Text() 返回当前行内容（不包含换行符）
		line := scanner.Text()
		fmt.Println(line)
	}
	// 检查扫描是否出错（不是 io.EOF）
	return scanner.Err()
}

// readBinaryFile 按块读取二进制文件
// 适合处理图片、视频等二进制文件
func readBinaryFile(filename string) ([]byte, error) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建缓冲区，每次读取 1KB
	buffer := make([]byte, 1024)
	var data []byte

	for {
		// Read 返回实际读取的字节数
		n, err := file.Read(buffer)
		if n > 0 {
			// 追加到结果切片
			data = append(data, buffer[:n]...)
		}
		// io.EOF 表示文件结束，不是错误
		if err == io.EOF {
			break
		}
		if err != nil {
			return data, err
		}
	}
	return data, nil
}

// writeSmallFile 写入小文件
// 适合一次性写入少量内容，会覆盖已存在的文件
func writeSmallFile(filename, content string) error {
	// WriteFile 自动处理文件创建和权限
	// 0644: owner 可读写，其他人只读
	return os.WriteFile(filename, []byte(content), 0644)
}

// appendToFile 追加写入文件
// 适合写日志，在文件末尾添加内容
func appendToFile(filename, entry string) error {
	// OpenFile 支持多种模式组合
	// O_APPEND: 追加模式
	// O_CREATE: 不存在则创建
	// O_WRONLY: 只写模式
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入字符串（内部会转换为字节）
	_, err = file.WriteString(entry)
	return err
}

// writeLargeFile 使用 bufio.Writer 批量写入
// 适合写入大量数据，缓冲区减少系统调用次数
func writeLargeFile(filename string, lines int) error {
	// Create 创建新文件（如果已存在会清空）
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建带缓冲的 Writer，默认缓冲区 4KB
	writer := bufio.NewWriter(file)
	for i := 0; i < lines; i++ {
		// Fprintf 格式化写入
		fmt.Fprintf(writer, "Line %d: data\n", i)
	}
	// Flush 将缓冲区内容写入文件
	// 忘记调用 Flush 会导致数据丢失
	return writer.Flush()
}

// copyFile 复制文件
// 使用 io.Copy 高效复制，支持任何 io.Reader 和 io.Writer
func copyFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// io.Copy 自动处理分块复制
	_, err = io.Copy(dstFile, srcFile)
	return err
}

// listDir 列出目录内容
// 使用 os.ReadDir 获取目录条目，比 ioutil.ReadDir 更高效
func listDir(dirname string) error {
	// ReadDir 返回 []DirEntry，包含名称和基本信息
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		// Info() 获取文件的详细信息
		info, _ := entry.Info()
		if info.IsDir() {
			fmt.Printf("[DIR]  %s\n", entry.Name())
		} else {
			// Size() 返回文件大小（字节）
			fmt.Printf("[FILE] %s (%d bytes)\n", entry.Name(), info.Size())
		}
	}
	return nil
}

func main() {
	// 示例：读取配置文件
	content, err := readSmallFile("config.txt")
	if err != nil {
		fmt.Println("Error reading config:", err)
	} else {
		fmt.Println("Config loaded:", content)
	}

	// 示例：写入日志
	err = appendToFile("app.log", "Application started\n")
	if err != nil {
		fmt.Println("Error writing log:", err)
	}

	// 示例：列出目录
	_ = listDir(".")

	// 下面这些函数需要实际文件存在才能运行
	// _ = copyFile
	// _ = readBinaryFile
	// _ = writeLargeFile
	// _ = readLargeFileLineByLine
}
