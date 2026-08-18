package main

import (
	"fmt"
	"runtime"
)

const sampleAllocationSize = 100 << 20

func main() {
	printMemStats("启动")

	// 写入每个页面，避免示例只分配虚拟地址而没有形成可观察的物理页。
	allocation := make([]byte, sampleAllocationSize)
	for i := range allocation {
		allocation[i] = byte(i)
	}
	printMemStats("分配后")

	// 清除最后一个业务引用，再手动触发 GC，观察存活堆和累计 GC 次数的变化。
	allocation = nil
	runtime.GC()
	printMemStats("GC 后")
}

// printMemStats 输出当前堆分配和累计 GC 次数。
func printMemStats(label string) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("%s: Alloc=%d MiB HeapAlloc=%d MiB NumGC=%d\n",
		label,
		stats.Alloc/(1024*1024),
		stats.HeapAlloc/(1024*1024),
		stats.NumGC,
	)
}
