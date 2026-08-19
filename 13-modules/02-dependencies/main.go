package main

import (
	"fmt"

	"github.com/fatih/color"
)

// 依赖管理示例

func main() {
	fmt.Println("=== 依赖管理示例 ===")
	fmt.Println()

	// 使用第三方库
	color.Green("成功添加依赖")
	color.Yellow("运行 go mod tidy 后需要检查依赖变化")
	color.Cyan("查看 go.mod 和 go.sum 文件")
	fmt.Println()

	fmt.Println("常用命令:")
	fmt.Println("  go get github.com/fatih/color@latest  # 添加最新版本")
	fmt.Println("  go get github.com/fatih/color@v1.15.0 # 添加特定版本")
	fmt.Println("  go mod tidy                            # 整理依赖")
	fmt.Println("  go list -m all                         # 查看所有依赖")
	fmt.Println("  go mod why github.com/fatih/color      # 查看依赖原因")
}
