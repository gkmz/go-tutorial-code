package main

import "fmt"

// 基础示例：初始化 Go 模块

func main() {
	fmt.Println("=== Go Modules 基础示例 ===")
	fmt.Println()

	fmt.Println("1. 初始化模块:")
	fmt.Println("   go mod init github.com/username/myproject")
	fmt.Println()

	fmt.Println("2. 查看 go.mod 文件:")
	fmt.Println("   module github.com/username/myproject")
	fmt.Println("   go 1.25.1")
	fmt.Println()

	fmt.Println("3. 添加依赖:")
	fmt.Println("   import \"github.com/gin-gonic/gin\"")
	fmt.Println("   go mod tidy")
	fmt.Println()

	fmt.Println("4. 查看依赖:")
	fmt.Println("   go list -m all")
	fmt.Println()

	fmt.Println("Hello, Go Modules!")
}
