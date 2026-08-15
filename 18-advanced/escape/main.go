package main

import "fmt"

// NewMessage 返回局部变量地址，用于配合 go test -gcflags=-m 观察逃逸分析。
func NewMessage(text string) *string { return &text }

func main() { fmt.Println(*NewMessage("escape analysis")) }
