package main

// User 是逃逸分析示例使用的用户值。
type User struct {
	Name string
}

// NewUserValue 按值返回局部变量。
func NewUserValue(name string) User {
	return User{Name: name}
}

// NewUserPointer 返回局部变量的地址。
// 编译器会根据调用上下文决定该对象最终位于栈还是堆。
func NewUserPointer(name string) *User {
	user := User{Name: name}
	return &user // user escapes to heap in NewUserPointer
}

// MakeNumbers 返回动态长度切片，底层数组需要跨越函数边界继续存活。
func MakeNumbers(size int) []int {
	return make([]int, size) // escapes to heap
}

// NewCounter 返回捕获局部变量的闭包。
func NewCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

var retainedUser *User

// RetainUser 把指针保存到包级变量中，强制其指向的对象跨越调用边界。
func RetainUser(user *User) {
	retainedUser = user
}
