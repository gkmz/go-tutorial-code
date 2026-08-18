// Command fixed 演示使用互斥锁同步 Map 写入。
package main

import "sync"

func main() {
	values := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	for value := 1; value <= 2; value++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			values["answer"] = value
			mu.Unlock()
		}()
	}
	wg.Wait()
}
