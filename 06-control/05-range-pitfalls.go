package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 陷阱 1：值是拷贝
	fmt.Println("=== Pitfall 1: Value is a copy ===")
	people := []Person{
		{"Alice", 20},
		{"Bob", 25},
	}

	fmt.Println("Before (wrong way):")
	fmt.Println(people)

	// 错误：修改的是拷贝，不会影响原数据
	for _, p := range people {
		p.Age++ // 无效！
	}

	fmt.Println("After (wrong way):")
	fmt.Println(people) // 年龄没有变化

	// 正确方式 1：使用索引访问
	fmt.Println("\nCorrect way 1: Use index")
	for i := range people {
		people[i].Age++
	}
	fmt.Println(people)

	// 正确方式 2：使用指针切片
	fmt.Println("\nCorrect way 2: Use pointer slice")
	ptrPeople := []*Person{
		{"Charlie", 30},
		{"David", 35},
	}

	for _, p := range ptrPeople {
		p.Age++ // 有效，因为 p 是指针
	}

	for _, p := range ptrPeople {
		fmt.Printf("%s: %d\n", p.Name, p.Age)
	}

	// 陷阱 2：循环变量地址问题
	fmt.Println("\n=== Pitfall 2: Loop variable address ===")

	// 不支持 错误做法（针对 Go 1.22 之前的版本）
	// 在 Go 1.22 之前，v 在整个循环中只有一个实例，每次迭代只是改变它的值
	// 所以取地址 &v 得到的永远是同一个内存地址
	var ptrs []*int
	nums := []int{1, 2, 3}
	for _, v := range nums {
		ptrs = append(ptrs, &v) // 危险！在旧版本中，所有指针都指向同一个 v
	}

	fmt.Println("Old way (Go < 1.22, all pointers point to the same address):")
	for i, ptr := range ptrs {
		fmt.Printf("ptrs[%d] = %d\n", i, *ptr)
	}
	// 在 Go 1.22 之前，输出都是 3
	// 在 Go 1.22 及之后，Go 编译器已经将 v 的作用域改为每次迭代创建一个新变量
	// 所以即使不显式创建新变量，输出也会是 1, 2, 3

	// 已完成 兼容性写法（在旧版本中手动创建新变量）
	var correctPtrs []*int
	for _, v := range nums {
		v := v // 创建新变量，遮蔽循环变量。在 Go 1.22+ 已经不需要这一行了
		correctPtrs = append(correctPtrs, &v)
	}

	fmt.Println("\nCompatible way (Manual shadowing):")
	for i, ptr := range correctPtrs {
		fmt.Printf("correctPtrs[%d] = %d\n", i, *ptr)
	}

	// 陷阱 3：修改 map 的值
	fmt.Println("\n=== Pitfall 3: Modifying map values ===")
	personMap := map[string]Person{
		"alice": {"Alice", 20},
		"bob":   {"Bob", 25},
	}

	// 错误：不能直接修改 map 中的结构体字段
	// for _, p := range personMap {
	//     p.Age++ // 编译错误
	// }

	// 正确方式 1：使用 key 访问
	fmt.Println("Correct way 1: Access by key")
	for key := range personMap {
		p := personMap[key]
		p.Age++
		personMap[key] = p
	}

	for key, p := range personMap {
		fmt.Printf("%s: %d\n", key, p.Age)
	}

	// 正确方式 2：使用指针 map
	fmt.Println("\nCorrect way 2: Use pointer map")
	ptrMap := map[string]*Person{
		"charlie": {"Charlie", 30},
		"david":   {"David", 35},
	}

	for _, p := range ptrMap {
		p.Age++
	}

	for key, p := range ptrMap {
		fmt.Printf("%s: %d\n", key, p.Age)
	}

	// 陷阱 4：range 表达式只求值一次
	fmt.Println("\n=== Pitfall 4: Range expression evaluated once ===")
	slice := []int{1, 2, 3}
	for i, v := range slice {
		fmt.Printf("i=%d, v=%d\n", i, v)
		if i == 0 {
			slice = append(slice, 4, 5) // 不会影响当前循环
		}
	}
	fmt.Println("Final slice:", slice)
}
