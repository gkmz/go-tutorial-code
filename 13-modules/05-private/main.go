package main

import (
	"fmt"
	"os"
)

// 私有仓库配置示例

func main() {
	fmt.Println("=== 私有仓库配置示例 ===")
	fmt.Println()

	// 显示环境变量
	fmt.Println("当前环境变量:")
	fmt.Printf("  GOPRIVATE: %s\n", os.Getenv("GOPRIVATE"))
	fmt.Printf("  GOPROXY:   %s\n", os.Getenv("GOPROXY"))
	fmt.Printf("  GOSUMDB:   %s\n", os.Getenv("GOSUMDB"))
	fmt.Println()

	fmt.Println("配置步骤:")
	fmt.Println()

	fmt.Println("1. 设置 GOPRIVATE:")
	fmt.Println("   go env -w GOPRIVATE=github.com/yourcompany/*")
	fmt.Println()

	fmt.Println("2. 配置 Git 认证 (SSH):")
	fmt.Println("   git config --global url.\"git@github.com:\".insteadOf \"https://github.com/\"")
	fmt.Println()

	fmt.Println("3. 配置 Git 认证 (Token):")
	fmt.Println("   使用凭据管理器或 CI 密钥注入，避免把 Token 写入 URL 或命令历史")
	fmt.Println()

	fmt.Println("4. 使用 .netrc 文件（仅在权限为 0600 且不提交时）:")
	fmt.Println("   # ~/.netrc")
	fmt.Println("   machine github.com")
	fmt.Println("   login username")
	fmt.Println("   password token")
	fmt.Println()

	fmt.Println("5. 配置代理:")
	fmt.Println("   go env -w GOPROXY=https://goproxy.cn,direct")
	fmt.Println()

	fmt.Println("验证配置:")
	fmt.Println("   go get github.com/yourcompany/private-repo")
}
