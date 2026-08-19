//go:build !urfave

package main

import (
	"fmt"
	"strconv"

	"github.com/hankmor/calc/pkg/calculator"
)

// parseArguments 将命令行参数解析为两个数字和一个运算符。
func parseArguments(args []string) (float64, string, float64, error) {
	if len(args) != 3 {
		return 0, "", 0, fmt.Errorf("需要三个参数：<num1> <operator> <num2>")
	}
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("第一个数字格式错误: %w", err)
	}
	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("第二个数字格式错误: %w", err)
	}
	return a, args[1], b, nil
}

// calculate 根据运算符调用计算器核心，并统一返回不支持运算符的错误。
func calculate(calc *calculator.Calculator, a float64, operator string, b float64) (float64, error) {
	switch operator {
	case "+":
		return calc.Add(a, b), nil
	case "-":
		return calc.Subtract(a, b), nil
	case "*":
		return calc.Multiply(a, b), nil
	case "/":
		return calc.Divide(a, b)
	default:
		return 0, fmt.Errorf("不支持的运算符: %q", operator)
	}
}
