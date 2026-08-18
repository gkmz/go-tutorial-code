package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func readSmallFile(filename string, maxBytes int64) ([]byte, error) {
	if maxBytes < 0 {
		return nil, fmt.Errorf("read %q: max bytes must not be negative", filename)
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open %q: %w", filename, err)
	}
	defer file.Close()

	// 多读一个字节，用于区分“刚好达到上限”和“已经超过上限”。
	data, err := io.ReadAll(io.LimitReader(file, maxBytes+1))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", filename, err)
	}
	if int64(len(data)) > maxBytes {
		return nil, fmt.Errorf("read %q: file exceeds %d bytes", filename, maxBytes)
	}
	return data, nil
}

func scanLines(reader io.Reader, maxTokenBytes int, process func(string) error) error {
	if maxTokenBytes <= 0 {
		return errors.New("max token bytes must be positive")
	}
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, min(64*1024, maxTokenBytes)), maxTokenBytes)
	for scanner.Scan() {
		if err := process(scanner.Text()); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func writeLines(filename string, lines []string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create %q: %w", filename, err)
	}
	defer func() {
		if closeErr := file.Close(); err == nil && closeErr != nil {
			err = fmt.Errorf("close %q: %w", filename, closeErr)
		}
	}()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err = fmt.Fprintln(writer, line); err != nil {
			return fmt.Errorf("write %q: %w", filename, err)
		}
	}
	if err = writer.Flush(); err != nil {
		return fmt.Errorf("flush %q: %w", filename, err)
	}
	return nil
}

func copyFile(sourcePath, targetPath string) (err error) {
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
		return fmt.Errorf("copy data: %w", err)
	}
	return nil
}

func main() {
	path := filepath.Join(".", "config.txt")
	data, err := readSmallFile(path, 1<<20)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Println("read config:", err)
		}
		return
	}
	fmt.Printf("config bytes: %d\n", len(data))
}
