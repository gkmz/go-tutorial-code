//go:build urfave

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/hankmor/calc/pkg/calculator"
	"github.com/urfave/cli/v2"
)

func main() {
	calc := calculator.New()

	app := &cli.App{
		Name:    "calc",
		Usage:   "使用 urfave/cli 重构的专业计算器",
		Version: "v2.0.0",
		Authors: []*cli.Author{
			{
				Name:  "极客老墨",
				Email: "hankmo.x@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "add",
				Aliases:   []string{"a"},
				Usage:     "执行加法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return fmt.Errorf("必须提供两个数字进行加法运算")
					}
					a, err := strconv.ParseFloat(c.Args().Get(0), 64)
					if err != nil {
						return fmt.Errorf("第一个数字格式错误: %v", err)
					}
					b, err := strconv.ParseFloat(c.Args().Get(1), 64)
					if err != nil {
						return fmt.Errorf("第二个数字格式错误: %v", err)
					}
					color.Green("Result: %.2f + %.2f = %.2f", a, b, calc.Add(a, b))
					return nil
				},
			},
			{
				Name:      "sub",
				Aliases:   []string{"s"},
				Usage:     "执行减法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return fmt.Errorf("必须提供两个数字进行减法运算")
					}
					a, _ := strconv.ParseFloat(c.Args().Get(0), 64)
					b, _ := strconv.ParseFloat(c.Args().Get(1), 64)
					color.Green("Result: %.2f - %.2f = %.2f", a, b, calc.Subtract(a, b))
					return nil
				},
			},
			{
				Name:      "mul",
				Aliases:   []string{"m"},
				Usage:     "执行乘法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return fmt.Errorf("必须提供两个数字进行乘法运算")
					}
					a, _ := strconv.ParseFloat(c.Args().Get(0), 64)
					b, _ := strconv.ParseFloat(c.Args().Get(1), 64)
					color.Green("Result: %.2f * %.2f = %.2f", a, b, calc.Multiply(a, b))
					return nil
				},
			},
			{
				Name:      "div",
				Aliases:   []string{"d"},
				Usage:     "执行除法运算",
				ArgsUsage: "<num1> <num2>",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return fmt.Errorf("必须提供两个数字进行除法运算")
					}
					a, _ := strconv.ParseFloat(c.Args().Get(0), 64)
					b, _ := strconv.ParseFloat(c.Args().Get(1), 64)
					res, err := calc.Divide(a, b)
					if err != nil {
						return err
					}
					color.Green("Result: %.2f / %.2f = %.2f", a, b, res)
					return nil
				},
			},
			{
				Name:  "history",
				Usage: "查看计算历史记录",
				Action: func(c *cli.Context) error {
					history := calc.History()
					if len(history) == 0 {
						color.Yellow("暂无计算记录")
						return nil
					}
					color.Cyan("计算历史:")
					for _, h := range history {
						fmt.Println("  ", h)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}
