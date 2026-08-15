package main

import (
	"fmt"
	"unsafe"
)

// SlicePitfalls 演示切片常见陷阱
func main() {
	fmt.Println("--- 陷阱 1: 切片扩容导致的底层数组变更 ---")
	s := make([]int, 0, 3)
	s = append(s, 1, 2, 3)

	// 记录原始底层数组地址
	originalPtr := unsafe.Pointer(&s[0])
	fmt.Printf("原始切片: %v, 地址: %p, 长度: %d, 容量: %d\n", s, originalPtr, len(s), cap(s))

	// 触发扩容
	s = append(s, 4)
	newPtr := unsafe.Pointer(&s[0])
	fmt.Printf("扩容后切片: %v, 地址: %p, 长度: %d, 容量: %d\n", s, newPtr, len(s), cap(s))

	if originalPtr != newPtr {
		fmt.Println("注意：底层数组地址已改变！")
	}

	fmt.Println("\n--- 陷阱 2: 多个切片共享底层数组的副作用 ---")
	base := []int{1, 2, 3, 4, 5}
	s1 := base[1:4] // [2, 3, 4]
	s2 := base[2:5] // [3, 4, 5]

	fmt.Printf("修改前 - s1: %v, s2: %v, base: %v\n", s1, s2, base)

	// 修改 s1 的一个元素，会影响 s2 和 base
	s1[1] = 99
	fmt.Printf("修改后 - s1: %v, s2: %v, base: %v\n", s1, s2, base)
	fmt.Println("注意：修改 s1[1] 导致 s2[0] 也变成了 99，因为它们共享底层内存！")

	fmt.Println("\n--- 避坑指南：如何安全复制切片 ---")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 100
	fmt.Printf("使用 copy 后 - src: %v, dst: %v (互不影响)\n", src, dst)
}
