package main

import (
	"fmt"
	"reflect"
)

// User 是反射示例使用的结构体。
type User struct {
	Name string
	Age  int
}

// Describe 返回结构体的字段名称和值，展示反射的基本边界。
func Describe(value any) []string {
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)
	result := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		result = append(result, fmt.Sprintf("%s=%v", typ.Field(i).Name, val.Field(i).Interface()))
	}
	return result
}

func main() { fmt.Println(Describe(User{Name: "老墨", Age: 18})) }
