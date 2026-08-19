package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// 添加全局 Flag 示例

var (
	verbose bool
	name    string
)

func main() {
	app := cli.NewApp()
	app.Name = "my-cli"
	app.Usage = "带全局 Flag 的 CLI 工具"
	app.Version = "0.1.0"

	// 全局 Flag
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name: "verbose",
			// -v 已被 urfave/cli 用作版本选项，详细模式使用 -V 避免冲突。
			Aliases:     []string{"V"},
			Usage:       "显示详细信息",
			Destination: &verbose,
		},
		&cli.StringFlag{
			Name:        "name",
			Aliases:     []string{"n"},
			Usage:       "你的名字",
			Value:       "World",
			Destination: &name,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		if verbose {
			fmt.Println("=== 详细模式 ===")
			fmt.Printf("Name: %s\n", name)
			fmt.Printf("Args: %v\n", ctx.Args().Slice())
		} else {
			fmt.Printf("Hello, %s!\n", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
