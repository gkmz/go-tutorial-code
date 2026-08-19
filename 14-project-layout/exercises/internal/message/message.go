// Package message 保存项目内部的业务逻辑。
package message

// Format 将名称格式化为命令行输出。
func Format(name string) string {
	if name == "" {
		name = "Go"
	}
	return "Hello, " + name + "!"
}

// Farewell 将名称格式化为命令行告别语。
func Farewell(name string) string {
	if name == "" {
		name = "Go"
	}
	return "Goodbye, " + name + "!"
}
