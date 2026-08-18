package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	// Block 和 mutex profile 默认关闭或采样较少，只在诊断期间按需启用。
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(10)

	// 让示例持续产生少量 CPU、分配和阻塞活动，便于观察 profile 结果。
	go runWorkload()

	// pprof 只应暴露在受保护的管理网络中，不要直接公开到公网。
	log.Fatal(http.ListenAndServe("127.0.0.1:6060", nil))
}

// runWorkload 生成可控的示例负载，不代表生产业务模型。
func runWorkload() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		data := make([]byte, 64*1024)
		for i := range data {
			data[i] = byte(i)
		}
		_ = data
	}
}
