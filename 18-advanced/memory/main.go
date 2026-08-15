package main

import (
	"fmt"
	"runtime"
)

func main() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("alloc=%d bytes heap=%d bytes\n", stats.Alloc, stats.HeapAlloc)
}
