package main

import "fmt"

// 常量示例
func constExamples() {
	fmt.Println("=== 常量 ===")

	// 单个常量
	const PI = 3.14159
	fmt.Printf("PI = %.5f\n", PI)

	// 批量常量
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)
	fmt.Printf("Status: %d, %d, %d\n", StatusOK, StatusNotFound, StatusError)

	// 常量组
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
		Sunday    = 7
	)
	fmt.Printf("Monday: %d, Friday: %d\n", Monday, Friday)
}

// iota 示例
func iotaExamples() {
	fmt.Println("\n=== iota ===")

	// iota 从 0 开始
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	// 简写（同一行省略）
	const (
		d = iota // 0
		e        // 1（继承 iota）
		f        // 2
	)
	fmt.Printf("d=%d, e=%d, f=%d\n", d, e, f)

	// 跳值用法
	const (
		_  = iota             // 0（跳过）
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20
		GB                    // 1 << 30
		TB                    // 1 << 40
	)
	fmt.Printf("KB=%d, MB=%d, GB=%d, TB=%d\n", KB, MB, GB, TB)

	// 位掩码
	const (
		FlagRead    = 1 << iota // 1
		FlagWrite               // 2
		FlagExecute             // 4
	)
	fmt.Printf("FlagRead=%d, FlagWrite=%d, FlagExecute=%d\n", FlagRead, FlagWrite, FlagExecute)

	// 混合使用
	const (
		CategoryBook        = iota // 0
		CategoryElectronics        // 1
		CategoryClothing           // 2
	)
	fmt.Printf("Categories: book=%d, electronics=%d, clothing=%d\n",
		CategoryBook, CategoryElectronics, CategoryClothing)
}

// 枚举类型示例
func enumExamples() {
	fmt.Println("\n=== 枚举类型 ===")

	// 使用 iota 定义枚举
	const (
		WeekdayMonday = iota
		WeekdayTuesday
		WeekdayWednesday
		WeekdayThursday
		WeekdayFriday
		WeekdaySaturday
		WeekdaySunday
	)

	// 使用枚举
	day := WeekdayFriday
	fmt.Printf("Today is: %d (Friday)\n", day)

	// 枚举判断
	if day == WeekdayFriday {
		fmt.Println("It's Friday! Weekend is coming!")
	}

	// 枚举列表
	days := []string{
		"Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday", "Sunday",
	}
	fmt.Printf("All days: %v\n", days)
}

func main() {
	constExamples()
	iotaExamples()
	enumExamples()
}
