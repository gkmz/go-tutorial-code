package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

// 整数类型示例
func integerTypes() {
	fmt.Println("=== 整数类型 ===")

	// 有符号整数
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n", i8, i16, i32, i64)

	// 无符号整数
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	fmt.Printf("uint8: %d, uint16: %d, uint32: %d, uint64: %d\n", u8, u16, u32, u64)

	// 平台相关类型
	var i int = 9223372036854775807
	var u uint = 18446744073709551615
	fmt.Printf("int (platform): %d, uint (platform): %d\n", i, u)

	// 特殊类型
	var b byte = 255 // uint8 别名
	var r rune = '中' // int32 别名，存 Unicode
	fmt.Printf("byte: %d, rune: %d (char: %c)\n", b, r, r)
}

// 浮点数类型示例
func floatTypes() {
	fmt.Println("\n=== 浮点数类型 ===")

	var f1 float32 = 3.1415926
	var f2 float64 = 3.141592653589793

	fmt.Printf("float32: %.7f\n", f1)
	fmt.Printf("float64: %.15f\n", f2)

	// 浮点数精度问题
	a := 0.1 + 0.2
	b := 0.3
	fmt.Printf("a = %.20f\n", a)
	fmt.Printf("b = %.20f\n", b)
	fmt.Printf("a == b: %v\n", a == b)

	// 正确比较方式
	diff := a - b
	if math.Abs(diff) < 1e-10 {
		fmt.Println("a and b are approximately equal")
	}
}

// 布尔类型示例
func boolType() {
	fmt.Println("\n=== 布尔类型 ===")

	var flag bool // 默认 false
	fmt.Printf("default bool: %v\n", flag)

	flag = true
	fmt.Printf("after assignment: %v\n", flag)

	// 逻辑运算
	fmt.Printf("true && false = %v\n", true && false)
	fmt.Printf("true || false = %v\n", true || false)
	fmt.Printf("!true = %v\n", !true)
}

// 字符串类型示例
func stringType() {
	fmt.Println("\n=== 字符串类型 ===")

	var s1 string // 默认 ""
	fmt.Printf("default string: %q\n", s1)

	s1 = "Hello, Go!"
	fmt.Printf("string: %s, len: %d\n", s1, len(s1))

	// 字符串是不可变的
	// s1[0] = 'h'  // 不支持 编译错误

	// 遍历字符串（按字节）
	fmt.Print("bytes: ")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c", s1[i])
	}
	fmt.Println()

	// 遍历字符串（按字符/rune）
	fmt.Print("chars: ")
	for _, ch := range s1 {
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	// 字符串拼接
	s2 := " is "
	s3 := "great"
	s4 := s1 + s2 + s3
	fmt.Println("concatenated:", s4)
}

// byte 和 rune 示例
func byteAndRune() {
	fmt.Println("\n=== byte 和 rune ===")

	// byte 存 ASCII
	var b byte = 'A'
	fmt.Printf("byte 'A': value=%d, char=%c\n", b, b)

	// rune 存 Unicode
	var r rune = '中'
	fmt.Printf("rune '中': value=%d, char=%c\n", r, r)

	// 遍历中文字符串
	s := "Hello, 世界"
	fmt.Printf("string: %s, len(bytes): %d\n", s, len(s))
	fmt.Printf("len(runes): %d\n", utf8.RuneCountInString(s))

	// 正确遍历
	fmt.Print("runes: ")
	for _, ch := range s {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}

func main() {
	integerTypes()
	floatTypes()
	boolType()
	stringType()
	byteAndRune()
}
