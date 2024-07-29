package main

import (
	"fmt"
)

func main() {
	app, err := InitializeApp("user:password@/dbname")
	if err != nil {
		fmt.Println("Failed to initialize app:", err)
		return
	}
	app.Run()
}
