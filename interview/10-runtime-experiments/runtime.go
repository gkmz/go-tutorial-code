// Package runtimeexperiments 提供 Runtime 与性能面试的最小实验入口。
package runtimeexperiments

import (
	"context"
	"io"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

// Snapshot 返回当前 Runtime 的基础指标快照。
type Snapshot struct {
	Goroutines  int
	HeapAlloc   uint64
	HeapObjects uint64
}

// CurrentSnapshot 读取当前 Goroutine 数和 Go Heap 指标。
func CurrentSnapshot() Snapshot {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return Snapshot{Goroutines: runtime.NumGoroutine(), HeapAlloc: stats.HeapAlloc, HeapObjects: stats.HeapObjects}
}

// EscapeToHeap 返回局部变量指针，用于配合 -gcflags=-m 观察逃逸分析。
func EscapeToHeap(value int) *int {
	return &value
}

// AllocateBytes 创建指定数量的临时字节对象，用于 Benchmark 和 Profile 实验。
func AllocateBytes(size int) []byte {
	if size < 0 {
		size = 0
	}
	return make([]byte, size)
}

// CaptureCPUProfile 采集指定时长的 CPU Profile。
func CaptureCPUProfile(ctx context.Context, output io.Writer, workload func()) error {
	if err := pprof.StartCPUProfile(output); err != nil {
		return err
	}
	defer pprof.StopCPUProfile()
	finished := make(chan struct{})
	go func() {
		workload()
		close(finished)
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-finished:
		return nil
	}
}

// CaptureTrace 将一段函数执行写入 trace 文件。
func CaptureTrace(path string, workload func()) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := trace.Start(file); err != nil {
		return err
	}
	workload()
	trace.Stop()
	return nil
}
