package main

import (
	"fmt"
	"sort"
)

// Person 用于演示自定义排序
type Person struct {
	Name string
	Age  int
}

// sortInts 整数排序
func sortInts() {
	nums := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	fmt.Println("Before:", nums)

	// Ints 原地排序 int 切片（升序）
	sort.Ints(nums)
	fmt.Println("After Ints:", nums)

	// 降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("Descending:", nums)
}

// sortStrings 字符串排序
func sortStrings() {
	strs := []string{"banana", "apple", "cherry", "date", "elderberry"}
	fmt.Println("Before:", strs)

	// Strings 原地排序字符串切片
	sort.Strings(strs)
	fmt.Println("After Strings:", strs)
}

// sortFloats 浮点数排序
func sortFloats() {
	nums := []float64{3.14, 1.41, 2.72, 0.58, 1.62}
	fmt.Println("Before:", nums)

	// Slice 通用排序，需要提供比较函数
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println("After Slice:", nums)
}

// sortCustom 自定义结构体排序
func sortCustom() {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Diana", Age: 28},
	}
	fmt.Println("Before:", people)

	// 按年龄排序（升序）
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("By Age (asc):", people)

	// 按姓名排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("By Name:", people)
}

// binarySearch 二分查找
func binarySearch() {
	nums := []int{1, 2, 5, 8, 9, 12, 15, 18, 20}
	fmt.Println("Array:", nums)

	// SearchInts 在已排序切片中查找
	// 返回第一个 >= target 的索引
	target := 8
	idx := sort.SearchInts(nums, target)
	if idx < len(nums) && nums[idx] == target {
		fmt.Printf("Found %d at index %d\n", target, idx)
	} else {
		fmt.Printf("%d not found, would be inserted at %d\n", target, idx)
	}

	// 查找不存在的值
	target = 10
	idx = sort.SearchInts(nums, target)
	fmt.Printf("%d not found, would be inserted at %d\n", target, idx)
}

// searchWithCondition 带条件的查找
func searchWithCondition() {
	nums := []int{1, 2, 5, 8, 9, 12, 15, 18, 20}
	fmt.Println("Array:", nums)

	// 查找第一个大于 10 的数
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] > 10
	})
	if idx < len(nums) {
		fmt.Printf("First number > 10 is %d at index %d\n", nums[idx], idx)
	}

	// 查找第一个偶数
	idx = sort.Search(len(nums), func(i int) bool {
		return nums[i]%2 == 0
	})
	if idx < len(nums) {
		fmt.Printf("First even number is %d at index %d\n", nums[idx], idx)
	}
}

// isSorted 检查是否已排序
func isSorted() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{5, 4, 3, 2, 1}

	fmt.Printf("%v is sorted: %v\n", nums1, sort.IntsAreSorted(nums1))
	fmt.Printf("%v is sorted: %v\n", nums2, sort.IntsAreSorted(nums2))
}

func main() {
	fmt.Println("=== sort.Ints ===")
	sortInts()

	fmt.Println("\n=== sort.Strings ===")
	sortStrings()

	fmt.Println("\n=== sort.Slice (floats) ===")
	sortFloats()

	fmt.Println("\n=== sort.Slice (custom) ===")
	sortCustom()

	fmt.Println("\n=== binary search ===")
	binarySearch()

	fmt.Println("\n=== search with condition ===")
	searchWithCondition()

	fmt.Println("\n=== is sorted check ===")
	isSorted()
}
