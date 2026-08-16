// Package exercises 提供文件与 IO 章节练习的参考实现。
package exercises

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// CopyFile 使用流式复制文件，适合不确定大小的输入。
func CopyFile(sourcePath, targetPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("open source: %w", err)
	}
	defer source.Close()
	target, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("create target: %w", err)
	}
	defer target.Close()
	if _, err := io.Copy(target, source); err != nil {
		return fmt.Errorf("copy file: %w", err)
	}
	return nil
}

// LowercaseTee 将输入中的 ASCII 大写字母转换为小写后写入目标和输出。
func LowercaseTee(source io.Reader, target io.Writer, output io.Writer) error {
	tee := io.TeeReader(source, lowercaseWriter{target: target})
	buffer := make([]byte, 32*1024)
	for {
		n, err := tee.Read(buffer)
		if n > 0 {
			converted := strings.ToLower(string(buffer[:n]))
			if _, writeErr := output.Write([]byte(converted)); writeErr != nil {
				return writeErr
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

type lowercaseWriter struct {
	target io.Writer
}

func (w lowercaseWriter) Write(data []byte) (int, error) {
	converted := strings.ToLower(string(data))
	if _, err := io.WriteString(w.target, converted); err != nil {
		return 0, err
	}
	return len(data), nil
}

// CountWords 使用 Scanner 统计空白分隔的单词。
func CountWords(source io.Reader) (map[string]int, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(source)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	return counts, scanner.Err()
}
