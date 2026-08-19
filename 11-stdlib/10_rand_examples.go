package main

import (
	"fmt"
	"math/rand"
	"time"
)

// randomInt 输出多个整数区间的随机样本。
func randomInt(random *rand.Rand) {
	// Intn(n) 生成 [0, n) 内的整数。
	fmt.Println("Random [0, 100):", random.Intn(100))
	fmt.Println("Random [1, 10]:", random.Intn(10)+1)
	fmt.Println("Random [-5, 5]:", random.Intn(11)-5)
}

// randomFloat 输出浮点数区间的随机样本。
func randomFloat(random *rand.Rand) {
	fmt.Println("Random [0.0, 1.0):", random.Float64())

	minValue, maxValue := 0.5, 2.5
	value := minValue + random.Float64()*(maxValue-minValue)
	fmt.Printf("Random [%.2f, %.2f): %.4f\n", minValue, maxValue, value)
}

// randomChoice 演示从非空切片中随机选择元素。
func randomChoice(random *rand.Rand) {
	colors := []string{"red", "green", "blue", "yellow", "purple"}
	choice := colors[random.Intn(len(colors))]
	fmt.Printf("Random choice from %v: %s\n", colors, choice)
}

// shuffleSlice 演示原地打乱切片。
func shuffleSlice(random *rand.Rand) {
	colors := []string{"red", "green", "blue", "yellow", "purple"}
	fmt.Println("Before shuffle:", colors)

	random.Shuffle(len(colors), func(i, j int) {
		colors[i], colors[j] = colors[j], colors[i]
	})
	fmt.Println("After shuffle:", colors)
}

// generateRandomArray 生成闭区间 [minValue, maxValue] 内的随机整数。
func generateRandomArray(random *rand.Rand, count, minValue, maxValue int) []int {
	if count <= 0 || minValue > maxValue {
		return []int{}
	}

	result := make([]int, count)
	for i := range result {
		result[i] = random.Intn(maxValue-minValue+1) + minValue
	}
	return result
}

// reproducibleSequence 演示固定种子产生可重复序列。
func reproducibleSequence(seed int64, count int) []int {
	random := rand.New(rand.NewSource(seed))
	return generateRandomArray(random, count, 0, 99)
}

func main() {
	// Go 1.25 中无需调用全局 rand.Seed；这里使用局部源明确状态所有权。
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("=== random int ===")
	randomInt(random)

	fmt.Println("\n=== random float ===")
	randomFloat(random)

	fmt.Println("\n=== random choice ===")
	randomChoice(random)

	fmt.Println("\n=== shuffle ===")
	shuffleSlice(random)

	fmt.Println("\n=== generate random array ===")
	fmt.Println(generateRandomArray(random, 10, 1, 100))

	fmt.Println("\n=== reproducible sequence ===")
	fmt.Println(reproducibleSequence(42, 5))
	fmt.Println(reproducibleSequence(42, 5))
}
