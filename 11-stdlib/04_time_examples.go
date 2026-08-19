package main

import (
	"fmt"
	"time"
)

// printTimeBasics 演示时间点、时区、格式化和时间间隔。
func printTimeBasics() {
	now := time.Now()
	fmt.Println("当前时间:", now)

	utc := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("UTC 时间:", utc)
	fmt.Println("Unix 时间戳:", now.Unix())
	fmt.Println("从时间戳恢复:", time.Unix(now.Unix(), 0))

	fmt.Println("日期:", now.Format("2006-01-02"))
	fmt.Println("日期时间:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("RFC3339:", now.Format(time.RFC3339))

	parsed, err := time.Parse("2006-01-02", "2024-01-12")
	if err != nil {
		fmt.Println("解析错误:", err)
	} else {
		fmt.Println("按 UTC 解析:", parsed)
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("加载时区错误:", err)
	} else {
		local, parseErr := time.ParseInLocation(
			"2006-01-02 15:04",
			"2026-08-19 09:30",
			location,
		)
		if parseErr != nil {
			fmt.Println("按指定时区解析错误:", parseErr)
		} else {
			fmt.Println("上海时间对应的 UTC:", local.UTC().Format(time.RFC3339))
		}
	}

	later := now.Add(2 * time.Hour)
	earlier := now.Add(-30 * time.Minute)
	fmt.Println("时间差:", later.Sub(now))
	fmt.Println("later 在 now 之后:", later.After(now))
	fmt.Println("earlier 在 now 之前:", earlier.Before(now))
	fmt.Println("常用间隔:", time.Second, time.Minute, time.Hour)
}

// timerExample 演示一次性 Timer 如何与业务结果竞争。
func timerExample() {
	fmt.Println("\n=== Timer：一次性事件 ===")
	timer := time.NewTimer(30 * time.Millisecond)
	defer timer.Stop()

	resultCh := make(chan string, 1)
	select {
	case result := <-resultCh:
		fmt.Println("业务结果:", result)
	case <-timer.C:
		fmt.Println("Timer 到期：操作超时")
	}
}

// tickerExample 演示 Ticker 如何产生周期事件并在退出时停止。
func tickerExample() {
	fmt.Println("\n=== Ticker：周期事件 ===")
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	for index := 1; index <= 3; index++ {
		<-ticker.C
		fmt.Println("第", index, "次周期事件")
	}
	fmt.Println("Ticker 已停止")
}

func main() {
	fmt.Println("=== 时间点、时区与间隔 ===")
	printTimeBasics()
	timerExample()
	tickerExample()
}
