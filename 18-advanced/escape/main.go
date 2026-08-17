package main

import "fmt"

func main() {
	user := NewUserPointer("Hank")
	numbers := MakeNumbers(4)
	counter := NewCounter()
	fmt.Println(user.Name, numbers, counter(), counter())
}
