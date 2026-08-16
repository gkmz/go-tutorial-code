# 练习 5 参考答案

原程序可以运行，但有以下可改进点：

1. `GetUserName` 是导出函数，正式项目中应补充以 `GetUserName` 开头的文档注释。
2. `if ok == true` 可以简化为 `if ok`，直接表达布尔判断。
3. `else` 紧跟在 `return` 后面没有必要，可以提前返回空字符串，减少嵌套。
4. `var name string = GetUserName(1)` 可以简化为 `name := GetUserName(1)`。
5. `os` 和 `strings` 被导入但没有使用，会导致编译失败，应删除。
6. `users` 每次调用都重新创建。示例规模很小没有实际性能问题，但生产代码应根据生命周期决定是否复用数据。

一种修改结果如下：

```go
package main

import "fmt"

// GetUserName 根据用户 ID 返回用户名；找不到时返回空字符串。
func GetUserName(id int64) string {
	users := map[int64]string{
		1: "zhangsan",
		2: "lisi",
	}
	name, ok := users[id]
	if !ok {
		return ""
	}
	return name
}

func main() {
	name := GetUserName(1)
	fmt.Println(name)
}
```
