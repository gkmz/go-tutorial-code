package main

import (
	"fmt"
	"os"

	"example.com/go-tutorial/14-project-layout-exercises/internal/message"
)

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(message.Farewell(name))
}
