package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// compareBytes 比较字节切片
func compareBytes() {
	a := []byte("hello")
	b := []byte("world")
	c := []byte("hello")

	// Compare 返回 -1 (a < b), 0 (a == b), 1 (a > b)
	fmt.Println("Compare(a, b):", bytes.Compare(a, b)) // -1
	fmt.Println("Compare(a, c):", bytes.Compare(a, c)) // 0
	fmt.Println("Compare(b, a):", bytes.Compare(b, a)) // 1

	// Equal 检查是否相等
	fmt.Println("Equal(a, c):", bytes.Equal(a, c)) // true
	fmt.Println("Equal(a, b):", bytes.Equal(a, b)) // false
}

// containsBytes 检查字节切片是否包含子切片
func containsBytes() {
	data := []byte("hello world, hello Go!")

	// Contains 检查是否包含
	fmt.Printf("Contains 'world': %v\n", bytes.Contains(data, []byte("world")))
	fmt.Printf("Contains 'Go': %v\n", bytes.Contains(data, []byte("Go")))
	fmt.Printf("Contains 'Rust': %v\n", bytes.Contains(data, []byte("Rust")))

	// ContainsAny 检查是否包含任意字符
	fmt.Printf("ContainsAny 'aeiou': %v\n", bytes.ContainsAny(data, "aeiou"))
	fmt.Printf("ContainsAny 'xyz': %v\n", bytes.ContainsAny(data, "xyz"))
}

// countBytes 统计子切片出现次数
func countBytes() {
	data := []byte("hello world, hello again!")

	// Count 统计子切片出现次数
	count := bytes.Count(data, []byte("l"))
	fmt.Printf("Count of 'l' in %q: %d\n", string(data), count)

	// 统计空格
	spaceCount := bytes.Count(data, []byte(" "))
	fmt.Printf("Count of space: %d\n", spaceCount)

	// 统计 "hello"
	helloCount := bytes.Count(data, []byte("hello"))
	fmt.Printf("Count of 'hello': %d\n", helloCount)
}

// splitBytes 分割字节切片
func splitBytes() {
	data := []byte("a,b,c,d,e")

	// Split 按分隔符分割
	parts := bytes.Split(data, []byte(","))
	fmt.Printf("Split by ',': %v\n", parts)

	// SplitN 限制分割数量
	parts2 := bytes.SplitN(data, []byte(","), 3)
	fmt.Printf("SplitN by ',' (n=3): %v\n", parts2)

	// Fields 按空白字符分割
	data2 := []byte("  hello   world  ")
	fields := bytes.Fields(data2)
	fmt.Printf("Fields: %v\n", fields)

	// FieldsFunc 按自定义函数分割
	data3 := []byte("hello,world;test|string")
	fields3 := bytes.FieldsFunc(data3, func(r rune) bool {
		return r == ',' || r == ';' || r == '|'
	})
	fmt.Printf("FieldsFunc (sep , ; |): %v\n", fields3)
}

// joinBytes 合并字节切片
func joinBytes() {
	parts := [][]byte{[]byte("hello"), []byte("world"), []byte("Go")}

	// Join 合并字节切片
	result := bytes.Join(parts, []byte("-"))
	fmt.Printf("Join with '-': %s\n", string(result))

	result2 := bytes.Join(parts, []byte(" "))
	fmt.Printf("Join with ' ': %s\n", string(result2))

	result3 := bytes.Join(parts, []byte{})
	fmt.Printf("Join with '': %s\n", string(result3))
}

// replaceBytes 替换字节切片
func replaceBytes() {
	data := []byte("hello world, hello!")

	// Replace 替换指定次数
	replaced := bytes.Replace(data, []byte("hello"), []byte("Hi"), 1)
	fmt.Printf("Replace 'hello' -> 'Hi' (1 time): %s\n", string(replaced))

	// ReplaceAll 替换所有
	replaced2 := bytes.ReplaceAll(data, []byte("hello"), []byte("Hi"))
	fmt.Printf("ReplaceAll 'hello' -> 'Hi': %s\n", string(replaced2))

	// 替换单个字符
	data2 := []byte("hello")
	replaced3 := bytes.Replace(data2, []byte("l"), []byte("L"), 2)
	fmt.Printf("Replace 'l' -> 'L' (2 times): %s\n", string(replaced3))
}

// buildBytes 构建字节切片
func buildBytes() {
	// 使用 Buffer 高效构建
	var buf bytes.Buffer

	buf.WriteString("Hello")
	buf.WriteByte(',')
	buf.Write([]byte(" World!"))

	// 使用 Fprintf 格式化写入
	fmt.Fprintf(&buf, " You are using %s.", "Go")

	result := buf.Bytes()
	fmt.Printf("Built with Buffer: %s\n", string(result))

	// 使用 NewBuffer 从现有数据创建
	buf2 := bytes.NewBuffer([]byte("Initial data"))
	buf2.WriteString(" + appended")
	fmt.Printf("NewBuffer: %s\n", buf2.String())

	// 使用 Grow 预分配内存
	buf3 := new(bytes.Buffer)
	buf3.Grow(1024) // 预分配 1KB 内存
	for i := 0; i < 100; i++ {
		buf3.WriteString("x")
	}
	fmt.Printf("Buffer with Grow: %d bytes\n", buf3.Len())
}

// readBytes 读取字节数据
func readBytes() {
	data := []byte("hello world this is a test")

	// Reader 创建字节读取器
	reader := bytes.NewReader(data)

	// Read 读取数据
	buf := make([]byte, 5)
	n, _ := reader.Read(buf)
	fmt.Printf("Read %d bytes: %q\n", n, string(buf[:n]))

	// ReadByte 读取单个字节
	b, _ := reader.ReadByte()
	fmt.Printf("ReadByte: %q\n", string(b))

	// bufio.Reader 提供按分隔符读取的 ReadBytes 方法。
	buffered := bufio.NewReader(reader)
	rest, _ := buffered.ReadBytes(' ')
	fmt.Printf("ReadBytes ' ': %q\n", string(rest))

	// ReadString 读取直到分隔符（返回字符串）
	reader2 := bufio.NewReader(bytes.NewReader(data))
	s, _ := reader2.ReadString(' ')
	fmt.Printf("ReadString ' ': %q\n", s)
}

// mapBytes 对字节应用函数
func mapBytes() {
	data := []byte("Hello World!")

	// Map 对每个字节应用函数
	mapped := bytes.Map(func(r rune) rune {
		// 转小写
		if r >= 'A' && r <= 'Z' {
			return r + 32
		}
		return r
	}, data)
	fmt.Printf("To lower: %s\n", string(mapped))

	// 另一种用法：替换字符
	mapped2 := bytes.Map(func(r rune) rune {
		if r == ' ' {
			return '-'
		}
		return r
	}, data)
	fmt.Printf("Replace space with '-': %s\n", string(mapped2))
}

// bufferExamples Buffer 的其他用法
func bufferExamples() {
	// Buffer 作为 io.Reader
	reader := bytes.NewBuffer([]byte("test data"))

	// 使用 io.Copy
	var output bytes.Buffer
	io.Copy(&output, reader)
	fmt.Printf("io.Copy result: %s\n", output.String())

	// Buffer 的 String 方法
	buf := bytes.NewBuffer([]byte("hello"))
	buf.WriteString(" world")
	fmt.Printf("Buffer.String(): %s\n", buf.String())

	// Reset 清空 Buffer
	buf.Reset()
	buf.WriteString("new content")
	fmt.Printf("After Reset: %s\n", buf.String())

	// Truncate 截断
	buf2 := bytes.NewBuffer([]byte("hello world"))
	buf2.Truncate(5)
	fmt.Printf("After Truncate(5): %s\n", buf2.String())

	// Bytes 返回切片（不会复制）
	buf3 := bytes.NewBuffer([]byte("test"))
	slice := buf3.Bytes()
	fmt.Printf("Bytes slice: %v\n", slice)
}

func main() {
	fmt.Println("=== compare bytes ===")
	compareBytes()

	fmt.Println("\n=== contains bytes ===")
	containsBytes()

	fmt.Println("\n=== count bytes ===")
	countBytes()

	fmt.Println("\n=== split bytes ===")
	splitBytes()

	fmt.Println("\n=== join bytes ===")
	joinBytes()

	fmt.Println("\n=== replace bytes ===")
	replaceBytes()

	fmt.Println("\n=== build bytes ===")
	buildBytes()

	fmt.Println("\n=== read bytes ===")
	readBytes()

	fmt.Println("\n=== map bytes ===")
	mapBytes()

	fmt.Println("\n=== buffer examples ===")
	bufferExamples()
}
