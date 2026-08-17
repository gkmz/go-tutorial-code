package runtimeexperiments

import "testing"

// BenchmarkAllocateBytes 对比分配规模变化对分配次数和耗时的影响。
func BenchmarkAllocateBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AllocateBytes(1024)
	}
}
