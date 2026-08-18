package main

import (
	"os"
	"runtime/pprof"
)

// profileCPU 将指定函数执行期间的 CPU profile 写入文件。
func profileCPU(path string, run func()) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		return err
	}
	defer pprof.StopCPUProfile()

	run()
	return nil
}

// writeHeapProfile 将当前 Go 堆 profile 写入文件。
func writeHeapProfile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return pprof.WriteHeapProfile(file)
}
