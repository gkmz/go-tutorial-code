// Package exercises 提供可测试的 CLI 应用构造函数。
package exercises

import (
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

// ErrNegativeCount 表示 count 命令收到了负数。
var ErrNegativeCount = errors.New("count 不能为负数")

// NewApp 创建一个带 greet 子命令的 CLI 应用。
func NewApp() *cli.App {
	return &cli.App{
		Name: "greet",
		Commands: []*cli.Command{{
			Name: "hello",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Required: true},
				&cli.BoolFlag{Name: "upper", Usage: "将问候语转换为大写"},
			},
			Action: func(ctx *cli.Context) error {
				message := fmt.Sprintf("Hello, %s!\n", ctx.String("name"))
				if ctx.Bool("upper") {
					message = strings.ToUpper(message)
				}
				_, err := fmt.Fprint(ctx.App.Writer, message)
				return err
			},
		}, {
			Name: "count",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "number", Required: true, Usage: "要计数的非负整数"},
			},
			Action: func(ctx *cli.Context) error {
				number := ctx.Int("number")
				if number < 0 {
					return ErrNegativeCount
				}
				_, err := fmt.Fprintf(ctx.App.Writer, "Count: %d\n", number)
				return err
			},
		}},
	}
}
