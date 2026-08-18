package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProfileWriters(t *testing.T) {
	directory := t.TempDir()
	cpuPath := filepath.Join(directory, "cpu.pprof")
	heapPath := filepath.Join(directory, "heap.pprof")

	if err := profileCPU(cpuPath, func() {
		for i := 0; i < 10000; i++ {
			_ = i * i
		}
	}); err != nil {
		t.Fatalf("profileCPU() error = %v", err)
	}
	if err := writeHeapProfile(heapPath); err != nil {
		t.Fatalf("writeHeapProfile() error = %v", err)
	}

	for _, path := range []string{cpuPath, heapPath} {
		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("stat(%q) error = %v", path, err)
		}
		if info.Size() == 0 {
			t.Fatalf("profile %q is empty", path)
		}
	}
}
