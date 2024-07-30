package main

import (
	"fmt"
	"wire-example/ioc"
)

func main() {
	app, err := ioc.InitializeApp()
	if err != nil {
		fmt.Println("Failed to initialize app:", err)
		return
	}
	app.Run()
}
