package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("--- 场景 1: 原生 Map 的并发写冲突 (引发 panic) ---")

	// 模拟并发更新 Map
	// m := make(map[string]int)
	// 注意：运行此代码时，可能会直接引发：fatal error: concurrent map writes
	// go run -race 可检测此类问题
	/*
		for i := 0; i < 100; i++ {
			go func(val int) {
				m["key"] = val
			}(i)
		}
	*/
	fmt.Println("提示：直接并发写入原生 Map 会导致程序崩溃 (fatal error: concurrent map writes)")

	fmt.Println("\n--- 场景 2: 使用 sync.RWMutex 保证并发安全 ---")
	type SafeMap struct {
		sync.RWMutex
		data map[string]int
	}

	sm := &SafeMap{data: make(map[string]int)}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sm.Lock() // 写锁
			sm.data["key"] = val
			sm.Unlock()
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.RLock() // 读锁
			_ = sm.data["key"]
			sm.RUnlock()
		}()
	}

	wg.Wait()
	fmt.Println("成功使用 RWMutex 完成并发读写，未发生 panic")

	time.Sleep(100 * time.Millisecond)
}
