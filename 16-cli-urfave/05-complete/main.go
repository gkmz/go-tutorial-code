package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// 完整示例：包含全局 Flag、命令、子命令、生命周期方法

var (
	verbose  bool
	weathers = []string{"sunny", "windy", "cloudy", "rainy"}
)

func main() {
	app := cli.NewApp()
	app.Name = "demo-cli"
	app.Usage = "完整的 CLI 工具示例"
	app.Version = "1.0.0"

	// 全局 Flag
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"v"},
			Usage:       "显示详细信息",
			Destination: &verbose,
		},
	}

	// 全局 Before
	app.Before = func(ctx *cli.Context) error {
		if verbose {
			fmt.Println("=== App Before ===")
		}
		return nil
	}

	// 全局 After
	app.After = func(ctx *cli.Context) error {
		if verbose {
			fmt.Println("=== App After ===")
		}
		return nil
	}

	// 定义命令
	app.Commands = []*cli.Command{
		{
			Name:    "hello",
			Aliases: []string{"h"},
			Usage:   "向你问好",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "你的名字",
					Required: true,
				},
			},
			Before: func(ctx *cli.Context) error {
				if verbose {
					fmt.Println("=== Command Before ===")
				}
				return nil
			},
			After: func(ctx *cli.Context) error {
				if verbose {
					fmt.Println("=== Command After ===")
				}
				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:     "weather",
					Aliases:  []string{"w"},
					Usage:    "报告天气情况",
					Category: "info",
					Action: func(ctx *cli.Context) error {
						name := ctx.String("name")
						rd := rand.New(rand.NewSource(time.Now().UnixNano()))
						weather := weathers[rd.Intn(len(weathers))]
						fmt.Printf("Hello %s, today is a %s day!\n", name, weather)
						return nil
					},
				},
				{
					Name:     "time",
					Aliases:  []string{"t"},
					Usage:    "报告当前时间",
					Category: "info",
					Action: func(ctx *cli.Context) error {
						name := ctx.String("name")
						now := time.Now().Format("15:04:05")
						fmt.Printf("Hello %s, current time is %s\n", name, now)
						return nil
					},
				},
			},
			Action: func(ctx *cli.Context) error {
				name := ctx.String("name")
				fmt.Printf("Hello, %s!\n", name)
				return nil
			},
		},
		{
			Name:    "goodbye",
			Aliases: []string{"bye"},
			Usage:   "说再见",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "你的名字",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				name := ctx.String("name")
				fmt.Printf("Goodbye, %s! See you next time.\n", name)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
