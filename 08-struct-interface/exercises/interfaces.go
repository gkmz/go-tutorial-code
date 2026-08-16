// Package exercises 提供结构体与接口章节练习的参考实现。
package exercises

import (
	"encoding/json"
	"fmt"
	"math"
)

// Animal 定义动物的基本行为。
type Animal interface {
	Eat() string
	Sleep() string
}

// Dog 是 Animal 的一种实现。
type Dog struct{}

// Eat 返回狗的进食描述。
func (Dog) Eat() string { return "dog eats" }

// Sleep 返回狗的睡眠描述。
func (Dog) Sleep() string { return "dog sleeps" }

// Cat 是 Animal 的一种实现。
type Cat struct{}

// Eat 返回猫的进食描述。
func (Cat) Eat() string { return "cat eats" }

// Sleep 返回猫的睡眠描述。
func (Cat) Sleep() string { return "cat sleeps" }

// Address 表示地址信息。
type Address struct {
	City   string
	Street string
}

// Person 通过嵌入 Address 复用地址字段。
type Person struct {
	Name string
	Address
}

// Shape 定义形状的面积和周长行为。
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 表示矩形。
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 返回矩形面积。
func (r Rectangle) Area() float64 { return r.Width * r.Height }

// Perimeter 返回矩形周长。
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Circle 表示圆形。
type Circle struct{ Radius float64 }

// Area 返回圆形面积。
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// Perimeter 返回圆形周长。
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// TotalArea 返回一组形状的总面积。
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// DescribeType 返回类型 switch 的类型名称。
func DescribeType(value any) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		return "float64"
	default:
		return "unknown"
	}
}

// TypedNilError 演示带 nil 指针接收者的错误类型。
type TypedNilError struct{}

// Error 返回错误描述。
func (*TypedNilError) Error() string { return "typed nil error" }

// WrongError 返回一个非 nil 的 typed nil error 接口。
func WrongError() error {
	var err *TypedNilError
	return err
}

// CorrectError 返回真正的 nil error 接口。
func CorrectError() error { return nil }

// MarshalUser 将用户结构体序列化为带标签的 JSON。
func MarshalUser() ([]byte, error) {
	user := struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Password string `json:"-"`
	}{1, "Go", "secret"}
	return json.Marshal(user)
}

// FormatAnimal 返回动物行为摘要。
func FormatAnimal(animal Animal) string {
	return fmt.Sprintf("%s; %s", animal.Eat(), animal.Sleep())
}
