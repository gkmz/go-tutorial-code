// Package greeter 提供 Go workspace 示例使用的问候能力。
package greeter

// Message 返回包含指定名称的问候语。
func Message(name string) string {
	if name == "" {
		name = "Go"
	}
	return "Hello, " + name + "!"
}
