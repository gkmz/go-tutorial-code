//go:build cgo

package main

import "fmt"

func main() {
	message, err := duplicateMessage("hello from Go")
	if err != nil {
		fmt.Println("duplicate message:", err)
		return
	}

	// CGO 会引入跨语言调用成本，指针传递还必须遵守 CGO 指针规则。
	fmt.Println("add:", add(2, 3))
	fmt.Println("message:", message)
}
