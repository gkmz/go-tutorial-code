package main

import "fmt"

func main() {
	user := NewUserPointer("Hank")
	numbers := MakeNumbers(4)
	counter := NewCounter()
	fmt.Println(user.Name, numbers, counter(), counter())
}

// 使用  go build -o /tmp/go-escape-example -gcflags='-m=2' . 编译查看逃逸分析的逃逸情况，例如 `moved to heap`、`escapes to heap` 等信息。
