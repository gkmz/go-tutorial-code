package main

import (
	"context"
	"fmt"
	"log"

	wireexample "github.com/hankmor/go-tutorial-code/18-advanced/wire"
)

func main() {
	app, cleanup, err := wireexample.InitializeApp(wireexample.Config{
		DSN:            "memory://demo",
		GreetingPrefix: "Hello",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	message, err := app.Greet(context.Background(), "Wire")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
