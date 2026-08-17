package main

import (
	"context"
	"errors"
	"runtime"
	"sync"
	"time"
)

const cpuChunkSize = 1_000_000

type burstResult struct {
	Count          int
	PeakGoroutines int
	HeapBefore     uint64
	HeapAtPeak     uint64
	Elapsed        time.Duration
}

type preemptionResult struct {
	ObserverTicks int
	Checksum      uint64
	Elapsed       time.Duration
}

type timerResult struct {
	Checksum uint64
	Elapsed  time.Duration
}

func runGoroutineBurst(count int) (burstResult, error) {
	if count <= 0 {
		return burstResult{}, errors.New("count must be greater than zero")
	}

	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	// release 让所有任务在同一位置等待，避免前面的 Goroutine 在统计前已经退出。
	release := make(chan struct{})
	var ready sync.WaitGroup
	var done sync.WaitGroup
	ready.Add(count)

	started := time.Now()
	for range count {
		done.Go(func() {
			ready.Done()
			<-release
		})
	}
	ready.Wait()

	var peak runtime.MemStats
	runtime.ReadMemStats(&peak)
	result := burstResult{
		Count:          count,
		PeakGoroutines: runtime.NumGoroutine(),
		HeapBefore:     before.HeapAlloc,
		HeapAtPeak:     peak.HeapAlloc,
		Elapsed:        time.Since(started),
	}

	close(release)
	done.Wait()
	return result, nil
}

func runPreemptionExperiment(duration time.Duration) (preemptionResult, error) {
	if duration <= 0 {
		return preemptionResult{}, errors.New("duration must be greater than zero")
	}

	// 使用一个 P 可以排除“观察任务在另一个 P 上并行运行”的情况。
	previous := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(previous)

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	checksumCh := make(chan uint64, 1)
	go func() {
		checksumCh <- runCancelableCPU(ctx, cpuChunkSize)
	}()

	started := time.Now()
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	result := preemptionResult{}
	for {
		select {
		case <-ticker.C:
			result.ObserverTicks++
		case checksum := <-checksumCh:
			result.Checksum = checksum
			result.Elapsed = time.Since(started)
			return result, nil
		}
	}
}

func runCancelableCPU(ctx context.Context, chunkSize int) uint64 {
	if chunkSize < 1 {
		chunkSize = 1
	}

	var checksum uint64
	for {
		// 块内没有显式让出操作，用于观察运行时抢占；块间检查 Context 用于业务退出。
		for i := 0; i < chunkSize; i++ {
			checksum = checksum*33 + uint64(i)
		}
		select {
		case <-ctx.Done():
			return checksum
		default:
		}
	}
}

func runTimerExperiment(delay time.Duration) (timerResult, error) {
	if delay <= 0 {
		return timerResult{}, errors.New("delay must be greater than zero")
	}

	started := time.Now()
	resultCh := make(chan uint64, 1)
	go func() {
		resultCh <- cpuWork(2_000_000)
	}()

	// 当前 Goroutine 等待计时器时处于 waiting 状态，P 可以继续运行计算任务。
	timer := time.NewTimer(delay)
	defer timer.Stop()
	<-timer.C

	return timerResult{
		Checksum: <-resultCh,
		Elapsed:  time.Since(started),
	}, nil
}

func cpuWork(units int) uint64 {
	var checksum uint64
	for i := 0; i < units; i++ {
		checksum = checksum*33 + uint64(i)
	}
	return checksum
}
