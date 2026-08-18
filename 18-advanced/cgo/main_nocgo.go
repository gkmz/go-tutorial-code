//go:build !cgo

package main

import "fmt"

func main() {
	// 为禁用 CGO 的构建提供明确提示，而不是留下“没有 Go 源文件”的错误。
	fmt.Println("this example requires CGO_ENABLED=1 and a C compiler")
}
