package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	mode := flag.String("mode", "burst", "实验模式：burst、preempt 或 timer")
	count := flag.Int("count", 10_000, "burst 模式创建的 Goroutine 数量")
	duration := flag.Duration("duration", 500*time.Millisecond, "preempt 或 timer 模式的实验时长")
	flag.Parse()

	fmt.Printf("Go=%s GOMAXPROCS=%d NumCPU=%d\n", runtime.Version(), runtime.GOMAXPROCS(0), runtime.NumCPU())

	var err error
	switch *mode {
	case "burst":
		err = printBurstResult(*count)
	case "preempt":
		err = printPreemptionResult(*duration)
	case "timer":
		err = printTimerResult(*duration)
	default:
		err = fmt.Errorf("unknown mode %q", *mode)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "experiment failed:", err)
		os.Exit(1)
	}
}

func printBurstResult(count int) error {
	result, err := runGoroutineBurst(count)
	if err != nil {
		return err
	}
	fmt.Printf(
		"created=%d peak_goroutines=%d heap_before=%d heap_at_peak=%d elapsed=%s\n",
		result.Count,
		result.PeakGoroutines,
		result.HeapBefore,
		result.HeapAtPeak,
		result.Elapsed,
	)
	return nil
}

func printPreemptionResult(duration time.Duration) error {
	result, err := runPreemptionExperiment(duration)
	if err != nil {
		return err
	}
	fmt.Printf("duration=%s observer_ticks=%d checksum=%d\n", result.Elapsed, result.ObserverTicks, result.Checksum)
	return nil
}

func printTimerResult(duration time.Duration) error {
	result, err := runTimerExperiment(duration)
	if err != nil {
		return err
	}
	fmt.Printf("waited=%s checksum=%d\n", result.Elapsed, result.Checksum)
	return nil
}
