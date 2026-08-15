package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group
var calls atomic.Int32

// loadOnce 合并同一 key 的并发请求，避免缓存击穿时重复访问下游。
func loadOnce(key string) (string, error, bool) {
	value, err, shared := group.Do(key, func() (any, error) {
		calls.Add(1)
		time.Sleep(10 * time.Millisecond)
		return "value:" + key, nil
	})
	return value.(string), err, shared
}

func main() {
	value, err, shared := loadOnce("user:1")
	fmt.Println(value, err, shared, calls.Load())
}
