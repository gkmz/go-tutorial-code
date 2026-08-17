package main

import (
	"errors"
	"fmt"
	"time"
)

// 实战建议 4：实现重试机制

// 简单的重试函数
func retry(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if i > 0 {
			fmt.Printf("Retrying in %v...\n", sleep)
			time.Sleep(sleep)
			sleep *= 2 // 指数退避
		}

		fmt.Printf("Attempt %d/%d\n", i+1, attempts)
		err = fn()
		if err == nil {
			return nil
		}

		fmt.Printf("  Failed: %v\n", err)
	}

	return fmt.Errorf("after %d attempts, last error: %w", attempts, err)
}

// 模拟可能失败的操作
var attemptCount = 0

func unreliableOperation() error {
	attemptCount++
	if attemptCount < 3 {
		return errors.New("temporary failure")
	}
	return nil
}

// 高级重试：支持自定义退避策略
type RetryConfig struct {
	MaxAttempts int
	InitialWait time.Duration
	MaxWait     time.Duration
	Multiplier  float64
}

func retryWithConfig(config RetryConfig, fn func() error) error {
	var err error
	wait := config.InitialWait

	for i := 0; i < config.MaxAttempts; i++ {
		if i > 0 {
			if wait > config.MaxWait {
				wait = config.MaxWait
			}
			fmt.Printf("Waiting %v before retry...\n", wait)
			time.Sleep(wait)
			wait = time.Duration(float64(wait) * config.Multiplier)
		}

		fmt.Printf("Attempt %d/%d\n", i+1, config.MaxAttempts)
		err = fn()
		if err == nil {
			return nil
		}

		fmt.Printf("  Failed: %v\n", err)
	}

	return fmt.Errorf("after %d attempts, last error: %w", config.MaxAttempts, err)
}

func main() {
	fmt.Println("=== 重试机制示例 ===")

	// 示例 1：简单重试
	fmt.Println("\n--- 简单重试 ---")
	attemptCount = 0
	err := retry(3, 100*time.Millisecond, unreliableOperation)
	if err != nil {
		fmt.Println("Final error:", err)
	} else {
		fmt.Println("Operation succeeded")
	}

	// 示例 2：高级重试配置
	fmt.Println("\n--- 高级重试配置 ---")
	attemptCount = 0
	config := RetryConfig{
		MaxAttempts: 4,
		InitialWait: 50 * time.Millisecond,
		MaxWait:     500 * time.Millisecond,
		Multiplier:  2.0,
	}
	err = retryWithConfig(config, unreliableOperation)
	if err != nil {
		fmt.Println("Final error:", err)
	} else {
		fmt.Println("Operation succeeded")
	}

	// 示例 3：永远失败的操作
	fmt.Println("\n--- 永远失败的操作 ---")
	err = retry(3, 50*time.Millisecond, func() error {
		return errors.New("permanent failure")
	})
	if err != nil {
		fmt.Println("Final error:", err)
	}
}
