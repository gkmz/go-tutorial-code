package main

import (
	"fmt"
	"runtime"
)

func main() {
	// GOMAXPROCS 返回当前运行时允许并行执行 Go 代码的 P 数量。
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("goroutines:", runtime.NumGoroutine())
}
