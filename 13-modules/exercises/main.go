// Package exercises 演示主模块通过 replace 使用本地模块。
package exercises

import "example.com/go-tutorial/13-modules-exercises/greeter"

// Greeting 返回由本地替换模块生成的问候语。
func Greeting(name string) string { return greeter.Message(name) }
