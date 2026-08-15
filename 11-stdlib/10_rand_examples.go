package main

import (
	"fmt"
	"math/rand"
	"time"
)

// initRand 初始化随机数种子
// 重要：必须在 main 或使用随机数之前调用
func initRand() {
	// 使用当前时间作为种子，确保每次运行结果不同
	// 如果不设置种子，默认种子是 1，每次运行结果相同
	rand.Seed(time.Now().UnixNano())
}

// randomInt 生成指定范围内的随机整数
func randomInt() {
	// rand.Intn(n) 生成 [0, n) 的随机整数
	// 要生成 [min, max]，需要 rand.Intn(max-min+1) + min
	fmt.Println("Random [0, 100):", rand.Intn(100))
	fmt.Println("Random [1, 10]:", rand.Intn(10)+1)
	fmt.Println("Random [-5, 5]:", rand.Intn(11)-5)
}

// randomFloat 生成指定范围内的随机浮点数
func randomFloat() {
	// rand.Float64() 生成 [0, 1) 的随机浮点数
	fmt.Println("Random [0.0, 1.0):", rand.Float64())

	// 要生成 [min, max]，使用公式
	min, max := 0.5, 2.5
	random := min + rand.Float64()*(max-min)
	fmt.Printf("Random [%.2f, %.2f): %.4f\n", min, max, random)
}

// randomChoice 从切片中随机选择一个元素
func randomChoice() {
	colors := []string{"red", "green", "blue", "yellow", "purple"}
	choice := colors[rand.Intn(len(colors))]
	fmt.Printf("Random choice from %v: %s\n", colors, choice)

	// 从数字中选择
	numbers := []int{10, 20, 30, 40, 50}
	numChoice := numbers[rand.Intn(len(numbers))]
	fmt.Printf("Random number from %v: %d\n", numbers, numChoice)
}

// shuffleSlice 打乱切片顺序
func shuffleSlice() {
	colors := []string{"red", "green", "blue", "yellow", "purple"}
	fmt.Println("Before shuffle:", colors)

	// rand.Shuffle 打乱切片
	// 第一个参数是长度，第二个是交换函数
	rand.Shuffle(len(colors), func(i, j int) {
		colors[i], colors[j] = colors[j], colors[i]
	})
	fmt.Println("After shuffle:", colors)
}

// generateRandomArray 生成随机数组
func generateRandomArray(n, min, max int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}
	return arr
}

// randomPassword 生成随机密码
func randomPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// randomWithSource 使用自定义 Source
func randomWithSource() {
	// 创建自定义 Source（更安全，可以并发使用）
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	fmt.Println("Using custom source:")
	fmt.Println("Random int:", r.Intn(100))
	fmt.Println("Random float:", r.Float64())
}

func main() {
	// 初始化随机数
	initRand()

	fmt.Println("=== random int ===")
	randomInt()

	fmt.Println("\n=== random float ===")
	randomFloat()

	fmt.Println("\n=== random choice ===")
	randomChoice()

	fmt.Println("\n=== shuffle ===")
	shuffleSlice()

	fmt.Println("\n=== generate random array ===")
	arr := generateRandomArray(10, 1, 100)
	fmt.Println("Random array:", arr)

	fmt.Println("\n=== random password ===")
	password := randomPassword(16)
	fmt.Printf("Random password: %s\n", password)

	fmt.Println("\n=== random with custom source ===")
	randomWithSource()
}
