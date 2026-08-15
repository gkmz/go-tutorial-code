package main

import "fmt"

// iota 基础用法
func iotaBasics() {
	fmt.Println("=== iota 基础 ===")

	// iota 从 0 开始，每行 +1
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	// 简写（继承上一个表达式）
	const (
		d = iota // 0
		e        // 1（继承 d = iota）
		f        // 2（继承 e）
	)
	fmt.Printf("d=%d, e=%d, f=%d\n", d, e, f)

	// 新的 const 块，iota 重置
	const g = iota // 0
	const h = iota // 0（新的 const 块）
	fmt.Printf("g=%d, h=%d\n", g, h)
}

// 同一行多个常量
func sameLine() {
	fmt.Println("\n=== 同一行多个常量 ===")

	const (
		e, f = iota, iota + 10 // 0, 10（同一行，iota 都是 0）
		g, h                   // 1, 11（iota 变成 1）
	)
	fmt.Printf("e=%d, f=%d, g=%d, h=%d\n", e, f, g, h)

	// 混合使用
	const (
		x, y = iota, iota * 2 // 0, 0
		z, _ = iota, iota * 2 // 1, 2
		w, _ = iota, iota * 2 // 2, 4
	)
	fmt.Printf("x=%d, y=%d, z=%d, w=%d\n", x, y, z, w)
}

// iota 跳值
func iotaSkip() {
	fmt.Println("\n=== iota 跳值 ===")

	// 用 _ 占位
	const (
		_ = iota // 0（跳过）
		A        // 1
		B        // 2
		_        // 3（跳过）
		C        // 4
	)
	fmt.Printf("A=%d, B=%d, C=%d\n", A, B, C)

	// 显式指定值
	const (
		StatusPending = iota // 0
		StatusRunning        // 1
		_                    // 2（跳过）
		StatusSuccess        // 3
		StatusFailed         // 4
	)
	fmt.Printf("StatusPending=%d, StatusRunning=%d, StatusSuccess=%d, StatusFailed=%d\n",
		StatusPending, StatusRunning, StatusSuccess, StatusFailed)
}

// iota 表达式
func iotaExpression() {
	fmt.Println("\n=== iota 表达式 ===")

	// 位移
	const (
		Bit0 = 1 << iota // 1
		Bit1             // 2
		Bit2             // 4
		Bit3             // 8
	)
	fmt.Printf("Bit0=%d, Bit1=%d, Bit2=%d, Bit3=%d\n", Bit0, Bit1, Bit2, Bit3)

	// 乘积
	const (
		KB = 1 << (10 * iota) // 1024
		MB                    // 1048576
		GB                    // 1073741824
		TB                    // 1099511627776
	)
	fmt.Printf("KB=%d, MB=%d, GB=%d, TB=%d\n", KB, MB, GB, TB)

	// 加法
	const (
		Level1 = iota + 1 // 1
		Level2            // 2
		Level3            // 3
	)
	fmt.Printf("Level1=%d, Level2=%d, Level3=%d\n", Level1, Level2, Level3)
}

// iota 陷阱
func iotaTrap() {
	fmt.Println("\n=== iota 陷阱 ===")

	// 陷阱：表达式中的 iota
	const (
		a = iota * 2 // 0
		b            // 2（继承 a = iota * 2，iota=1）
		c            // 4（继承 b，iota=2）
		d = iota     // 3（重新使用 iota，iota=3）
		e            // 4（继承 d = iota，iota=4）
	)
	fmt.Printf("a=%d, b=%d, c=%d, d=%d, e=%d\n", a, b, c, d, e)
	fmt.Println("注意：d = iota 不是继承前面的表达式，而是使用当前的 iota 值")
}

// 月份枚举（从 1 开始）
func monthEnum() {
	fmt.Println("\n=== 月份枚举（从 1 开始） ===")

	// 方法 1：用 _ 占位
	const (
		_         = iota // 0（跳过）
		January          // 1
		February         // 2
		March            // 3
		April            // 4
		MayMonth         // 5
		June             // 6
		July             // 7
		August           // 8
		September        // 9
		October          // 10
		November         // 11
		December         // 12
	)
	fmt.Printf("January=%d, December=%d\n", January, December)

	// 方法 2：iota + 1
	const (
		Jan      = iota + 1 // 1
		Feb                 // 2
		Mar                 // 3
		Apr                 // 4
		MayShort            // 5
		Jun                 // 6
		Jul                 // 7
		Aug                 // 8
		Sep                 // 9
		Oct                 // 10
		Nov                 // 11
		Dec                 // 12
	)
	fmt.Printf("Jan=%d, Dec=%d\n", Jan, Dec)
}

func main() {
	iotaBasics()
	sameLine()
	iotaSkip()
	iotaExpression()
	iotaTrap()
	monthEnum()
}
