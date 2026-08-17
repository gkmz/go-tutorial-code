package main

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkCPUWorkGOMAXPROCS(b *testing.B) {
	values := []int{1, 2, runtime.NumCPU()}
	seen := make(map[int]struct{}, len(values))
	for _, procs := range values {
		if procs < 1 {
			continue
		}
		if _, ok := seen[procs]; ok {
			continue
		}
		seen[procs] = struct{}{}

		b.Run(fmt.Sprintf("P=%d", procs), func(b *testing.B) {
			previous := runtime.GOMAXPROCS(procs)
			defer runtime.GOMAXPROCS(previous)

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					cpuWork(20_000)
				}
			})
		})
	}
}
