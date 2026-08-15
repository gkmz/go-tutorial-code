package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker Pool 模式示例

// 任务结构
type Job struct {
	ID     int
	Number int
}

// 结果结构
type Result struct {
	Job    Job
	Result int
}

// Worker 函数
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(500 * time.Millisecond) // 模拟耗时操作

		// 计算结果
		result := Result{
			Job:    job,
			Result: job.Number * 2,
		}

		results <- result
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 启动 Worker Pool
	var wg sync.WaitGroup
	fmt.Printf("Starting %d workers\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 发送任务
	fmt.Printf("Sending %d jobs\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Number: j * 10}
	}
	close(jobs)

	// 等待所有 Worker 完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	fmt.Println("\nResults:")
	for result := range results {
		fmt.Printf("Job %d result: %d\n", result.Job.ID, result.Result)
	}

	fmt.Println("\nAll jobs completed")
}
