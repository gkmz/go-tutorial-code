// Command racy 故意执行未同步写入，用于观察 Race Detector 报告。
package main

import "sync"

func main() {
	values := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(2)
	for value := 1; value <= 2; value++ {
		go func() {
			defer wg.Done()
			// 这里故意不加锁；请使用 go run -race 观察报告。
			values["answer"] = value
		}()
	}
	wg.Wait()
}
