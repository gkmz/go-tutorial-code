// Package exercises 提供文件与 IO 章节练习的参考实现。
package exercises

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
)

// CopyFile 使用流式复制文件，适合不确定大小的输入。
func CopyFile(sourcePath, targetPath string) (err error) {
	source, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("open source: %w", err)
	}
	defer source.Close()
	target, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("create target: %w", err)
	}
	defer func() {
		if closeErr := target.Close(); err == nil && closeErr != nil {
			err = fmt.Errorf("close target: %w", closeErr)
		}
	}()
	if _, err = io.Copy(target, source); err != nil {
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
			converted := lowercaseASCII(buffer[:n])
			if writeErr := writeFull(output, converted); writeErr != nil {
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
	converted := lowercaseASCII(data)
	written, err := w.target.Write(converted)
	if err != nil {
		return written, err
	}
	if written != len(converted) {
		return written, io.ErrShortWrite
	}
	return len(data), nil
}

func writeFull(writer io.Writer, data []byte) error {
	for len(data) > 0 {
		written, err := writer.Write(data)
		if err != nil {
			return err
		}
		if written == 0 {
			return io.ErrShortWrite
		}
		data = data[written:]
	}
	return nil
}

func lowercaseASCII(data []byte) []byte {
	output := make([]byte, len(data))
	for index, value := range data {
		if value >= 'A' && value <= 'Z' {
			value += 'a' - 'A'
		}
		output[index] = value
	}
	return output
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

// CountLongWords 使用自定义 Scanner 缓冲读取长 token。
func CountLongWords(source io.Reader, maxToken int) (map[string]int, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(source)
	scanner.Buffer(make([]byte, 1024), maxToken)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	return counts, scanner.Err()
}

// ReadFSFile 从抽象文件系统中读取指定路径。
func ReadFSFile(fileSystem fs.FS, name string) ([]byte, error) {
	return fs.ReadFile(fileSystem, name)
}
