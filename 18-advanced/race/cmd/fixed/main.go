// Command fixed demonstrates synchronized map writes.
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
