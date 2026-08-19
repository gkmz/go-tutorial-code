//go:build !urfave

// main.go
package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hankmor/calc/pkg/calculator"
)

func main() {
	// 注意：Shell 会把 * 当作通配符，因此乘法运算符应写成 '*' 或 '\*'。
	a, operator, b, err := parseArguments(os.Args[1:])
	if err != nil {
		color.Red("Error: %v", err)
		color.Yellow("Usage: calc <num1> <operator> <num2>")
		os.Exit(1)
	}
	calc := calculator.New()
	result, err := calculate(calc, a, operator, b)
	if err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	// 彩色输出结果
	color.Green("Result: %.2f", result)

	// 显示历史记录
	if len(calc.History()) > 0 {
		color.Cyan("\nHistory:")
		for _, record := range calc.History() {
			fmt.Println("  " + record)
		}
	}
}
