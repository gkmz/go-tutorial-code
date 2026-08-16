// Package exercises 提供可测试的 CLI 应用构造函数。
package exercises

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// NewApp 创建一个带 greet 子命令的 CLI 应用。
func NewApp() *cli.App {
	return &cli.App{
		Name: "greet",
		Commands: []*cli.Command{{
			Name:  "hello",
			Flags: []cli.Flag{&cli.StringFlag{Name: "name", Required: true}},
			Action: func(ctx *cli.Context) error {
				_, err := fmt.Fprintf(ctx.App.Writer, "Hello, %s!\n", ctx.String("name"))
				return err
			},
		}},
	}
}
