package main

import (
	"fmt"
	"strings"
)

// strings 包示例

func main() {
	fmt.Println("=== 示例 1: 常用操作 ===")
	str := " Hello World "

	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Contains(str, "World"))
	fmt.Println(strings.HasPrefix(str, " Hello"))
	fmt.Println(strings.HasSuffix(str, "World "))
	fmt.Println(strings.Index(str, "World"))
	fmt.Println(strings.LastIndex(str, "o"))
	fmt.Println(strings.Count("hello", "l"))

	fmt.Println("\n=== 示例 2: 字符串转换 ===")
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.Replace("hello", "l", "L", 1))
	fmt.Println(strings.ReplaceAll("hello", "l", "L"))
	fmt.Println(strings.Repeat("Go", 3))

	fmt.Println("\n=== 示例 3: 分割和连接 ===")
	parts := strings.Split("a,b,c", ",")
	fmt.Println(parts)

	fields := strings.Fields("  hello  world  ")
	fmt.Println(fields)

	result := strings.Join([]string{"a", "b", "c"}, "-")
	fmt.Println(result)

	fmt.Println("\n=== 示例 4: 字符串构建器 ===")
	var builder strings.Builder

	for i := 0; i < 5; i++ {
		builder.WriteString("Go")
	}

	fmt.Println(builder.String())
}
