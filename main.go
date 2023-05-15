package main

import (
	"AccountApp/config"
	"fmt"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}

	// fmt.Println("Successfully connected to database!")
}
