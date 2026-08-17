package main

import "fmt"

func main() {
	user := User{Name: "Hank", Age: 18}
	fields, err := Describe(user)
	if err != nil {
		fmt.Println("describe user:", err)
		return
	}
	fmt.Println(fields)
}
