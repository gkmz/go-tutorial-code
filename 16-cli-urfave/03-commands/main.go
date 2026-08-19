package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// 添加命令示例

func main() {
	app := cli.NewApp()
	app.Name = "my-cli"
	app.Usage = "带命令的 CLI 工具"
	app.Version = "0.1.0"

	// 定义命令
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
				fmt.Printf("Goodbye, %s!\n", name)
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
