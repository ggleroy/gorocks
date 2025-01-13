package main

import (
	"fmt"
	"os"
)

func File() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensures file is closed when function exits

	// Perform file operations
	fmt.Println("File opened successfully")
}
