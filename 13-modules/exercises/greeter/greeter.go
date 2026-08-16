// Package greeter 提供模块替换练习使用的问候功能。
package greeter

// Message 返回指定名称的问候语。
func Message(name string) string {
	if name == "" {
		name = "Go"
	}
	return "Hello, " + name + "!"
}
