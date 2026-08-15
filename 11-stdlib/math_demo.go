package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"path/filepath"
	"sort"
	"time"
)

// ============ math 包 ============

// calculateCircleArea 计算圆的面积
func calculateCircleArea(radius float64) float64 {
	// math.Pi 是 π 常量
	// math.Pow 计算幂运算（radius 的 2 次方）
	return math.Pi * math.Pow(radius, 2)
}

// calculateDistance 计算两点之间的距离
func calculateDistance(x1, y1, x2, y2 float64) float64 {
	// math.Sqrt 计算平方根
	// math.Hypot 计算 sqrt(x*x + y*y)，更稳定
	dx := x2 - x1
	dy := y2 - y1
	return math.Hypot(dx, dy)
}

// formatNumber 格式化数字（保留小数、整数等）
func formatNumber(n float64) {
	// math.Ceil 向上取整
	fmt.Printf("Ceil(%.2f) = %.0f\n", n, math.Ceil(n))

	// math.Floor 向下取整
	fmt.Printf("Floor(%.2f) = %.0f\n", n, math.Floor(n))

	// math.Round 四舍五入
	fmt.Printf("Round(%.2f) = %.0f\n", n, math.Round(n))

	// math.Trunc 取整（截断小数部分）
	fmt.Printf("Trunc(%.2f) = %.0f\n", n, math.Trunc(n))
}

// calculateTrig 三角函数示例
func calculateTrig(angle float64) {
	// angle 是弧度，需要用 math.Radians 转换
	// 或者直接用弧度值
	rad := angle * math.Pi / 180

	// 三角函数
	fmt.Printf("Sin(%.0f°) = %.4f\n", angle, math.Sin(rad))
	fmt.Printf("Cos(%.0f°) = %.4f\n", angle, math.Cos(rad))
	fmt.Printf("Tan(%.0f°) = %.4f\n", angle, math.Tan(rad))

	// 反三角函数
	sinVal := 0.5
	fmt.Printf("Asin(%.1f) = %.2f°\n", sinVal, math.Asin(sinVal)*180/math.Pi)
}

// ============ math/rand 包 ============

// initRand 初始化随机数种子
// 重要：必须在 main 或使用随机数之前调用
func initRand() {
	// 使用当前时间作为种子，确保每次运行结果不同
	// 如果不设置种子，默认种子是 1，每次运行结果相同
	rand.Seed(time.Now().UnixNano())
}

// randomInt 生成指定范围内的随机整数
func randomInt(min, max int) int {
	// rand.Intn(n) 生成 [0, n) 的随机整数
	// 要生成 [min, max]，需要 rand.Intn(max-min+1) + min
	return rand.Intn(max-min+1) + min
}

// randomFloat 生成指定范围内的随机浮点数
func randomFloat(min, max float64) float64 {
	// rand.Float64() 生成 [0, 1) 的随机浮点数
	return min + rand.Float64()*(max-min)
}

// shuffleSlice 打乱切片顺序
func shuffleSlice(slice []string) {
	// rand.Shuffle 打乱切片
	// 第一个参数是长度，第二个是交换函数
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// randomChoice 从切片中随机选择一个元素
func randomChoice[T any](slice []T) T {
	index := rand.Intn(len(slice))
	return slice[index]
}

// ============ sort 包 ============

// sortInts 整数排序
func sortInts(nums []int) {
	// sort.Ints 原地排序 int 切片
	sort.Ints(nums)
}

// sortStrings 字符串排序
func sortStrings(strs []string) {
	// sort.Strings 原地排序字符串切片
	sort.Strings(strs)
}

// sortFloats 浮点数排序
func sortFloats(nums []float64) {
	// sort.Slice 通用排序，需要提供比较函数
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

// sortCustom 自定义结构体排序
type Person struct {
	Name string
	Age  int
}

func sortPeople(people []Person) {
	// 按年龄排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

// binarySearch 二分查找（要求数组已排序）
func binarySearch(nums []int, target int) int {
	// sort.Search 在已排序切片中查找
	// 返回第一个满足条件的索引
	idx := sort.SearchInts(nums, target)
	if idx < len(nums) && nums[idx] == target {
		return idx
	}
	return -1
}

// ============ path/filepath 包 ============

// joinPath 拼接路径
func joinPath(dir, file string) string {
	// filepath.Join 自动处理路径分隔符和斜杠
	// 自动处理不同操作系统的路径格式
	return filepath.Join(dir, "subdir", file)
}

// getFileInfo 获取文件信息
func getFileInfo(path string) {
	// filepath.Abs 获取绝对路径
	absPath, _ := filepath.Abs(path)
	fmt.Println("Absolute path:", absPath)

	// filepath.Base 获取文件名
	base := filepath.Base(path)
	fmt.Println("Base name:", base)

	// filepath.Dir 获取目录
	dir := filepath.Dir(path)
	fmt.Println("Directory:", dir)

	// filepath.Ext 获取扩展名
	ext := filepath.Ext(path)
	fmt.Println("Extension:", ext)

	// filepath.Split 分割路径和文件名
	dir2, base2 := filepath.Split(path)
	fmt.Println("Split - Dir:", dir2, "Base:", base2)
}

// cleanPath 清理路径
func cleanPath(path string) string {
	// filepath.Clean 清理路径（解析 . 和 ..）
	cleaned := filepath.Clean(path)
	fmt.Println("Cleaned path:", cleaned)
	return cleaned
}

// isAbsPath 判断是否为绝对路径
func checkPathType(path string) {
	fmt.Printf("IsAbs(%s) = %v\n", path, filepath.IsAbs(path))
}

// globMatch 路径匹配
func globMatch(pattern, path string) {
	// filepath.Match 路径匹配
	// * 匹配任意字符（不包括路径分隔符）
	// ? 匹配单个字符
	match, err := filepath.Match(pattern, path)
	if err != nil {
		fmt.Println("Match error:", err)
		return
	}
	fmt.Printf("Match(%s, %s) = %v\n", pattern, path, match)
}

// ============ bytes 包 ============

// compareBytes 比较字节切片
func compareBytes(a, b []byte) int {
	// bytes.Compare 比较两个字节切片
	// 返回 -1、0、1
	return bytes.Compare(a, b)
}

// containsBytes 检查字节切片是否包含子切片
func containsBytes(slice, sub []byte) bool {
	// bytes.Contains 检查是否包含
	return bytes.Contains(slice, sub)
}

// countBytes 统计子切片出现次数
func countBytes(slice, sub []byte) int {
	// bytes.Count 统计次数
	return bytes.Count(slice, sub)
}

// replaceBytes 替换字节切片
func replaceBytes(slice, old, new []byte, n int) []byte {
	// bytes.Replace 替换指定次数
	// n = -1 表示替换所有
	return bytes.Replace(slice, old, new, n)
}

// splitBytes 分割字节切片
func splitBytes(slice []byte, sep byte) [][]byte {
	// bytes.Split 分割字节切片
	return bytes.Split(slice, []byte{sep})
}

// joinBytes 合并字节切片
func joinBytes(slices [][]byte, sep []byte) []byte {
	// bytes.Join 合并字节切片
	return bytes.Join(slices, sep)
}

// buildBytes 构建字节切片
func buildBytes() []byte {
	// bytes.Buffer 高效构建字节数据
	var buf bytes.Buffer

	buf.WriteString("Hello")
	buf.WriteByte(',')
	buf.Write([]byte(" World!"))

	return buf.Bytes()
}

// ============ 主函数 ============

func main() {
	// 初始化随机数
	initRand()

	// math 示例
	fmt.Println("=== math ===")
	fmt.Printf("Circle area (r=5): %.2f\n", calculateCircleArea(5))
	fmt.Printf("Distance (0,0 to 3,4): %.2f\n", calculateDistance(0, 0, 3, 4))
	formatNumber(3.7)
	calculateTrig(30)

	// math/rand 示例
	fmt.Println("\n=== math/rand ===")
	fmt.Printf("Random int [1,100]: %d\n", randomInt(1, 100))
	fmt.Printf("Random float [0.0,1.0]: %.4f\n", randomFloat(0, 1))

	// sort 示例
	fmt.Println("\n=== sort ===")
	nums := []int{5, 2, 8, 1, 9}
	sortInts(nums)
	fmt.Println("Sorted ints:", nums)

	strs := []string{"banana", "apple", "cherry"}
	sortStrings(strs)
	fmt.Println("Sorted strings:", strs)

	people := []Person{{"Alice", 30}, {"Bob", 25}}
	sortPeople(people)
	fmt.Println("Sorted people:", people)

	// path/filepath 示例
	fmt.Println("\n=== path/filepath ===")
	joinPath("dir", "file.txt")
	getFileInfo("/home/user/docs/report.pdf")
	cleanPath("/home/user/../user/./docs")
	checkPathType("/home/user")
	checkPathType("docs")
	globMatch("*.txt", "report.txt")
	globMatch("file-?.txt", "file-1.txt")

	// bytes 示例
	fmt.Println("\n=== bytes ===")
	a := []byte("hello")
	b := []byte("world")
	fmt.Printf("Compare: %d\n", compareBytes(a, b))
	fmt.Printf("Contains 'lo': %v\n", containsBytes(a, []byte("lo")))
	fmt.Printf("Count 'l': %d\n", countBytes(a, []byte("l")))

	// 随机选择示例
	colors := []string{"red", "green", "blue"}
	fmt.Printf("Random color: %s\n", randomChoice(colors))
	shuffleSlice(colors)
	fmt.Println("Shuffled colors:", colors)
}
