package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// 子命令示例

var weathers = []string{"sunny", "windy", "cloudy", "rainy"}

func main() {
	app := cli.NewApp()
	app.Name = "my-cli"
	app.Usage = "带子命令的 CLI 工具"
	app.Version = "0.1.0"

	// 定义命令和子命令
	app.Commands = []*cli.Command{
		{
			Name: "hello",
			// h 已被内置 help 命令占用，使用 H 避免别名冲突。
			Aliases: []string{"H"},
			Usage:   "向你问好",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "你的名字",
					Required: true,
				},
			},
			// 子命令
			Subcommands: []*cli.Command{
				{
					Name:     "weather",
					Aliases:  []string{"w"},
					Usage:    "报告天气情况",
					Category: "weather",
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
					Category: "time",
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
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
