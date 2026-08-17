package main

import (
	"fmt"

	"github.com/username/local-package/utils"
)

// replace 指令示例

func main() {
	fmt.Println("=== Replace 指令示例 ===")
	fmt.Println()

	// 使用本地包
	result := utils.Add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)
	fmt.Println()

	fmt.Println("go.mod 中的 replace 指令:")
	fmt.Println("  replace github.com/username/local-package => ./local-package")
	fmt.Println()

	fmt.Println("使用场景:")
	fmt.Println("  1. 本地开发调试")
	fmt.Println("  2. 修改第三方库")
	fmt.Println("  3. 测试未发布的代码")
	fmt.Println()

	fmt.Println("注意  注意: 提交前记得删除 replace 指令")
}
