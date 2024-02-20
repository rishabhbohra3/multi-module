package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Print the current working directory
	fmt.Println("Current Working Directory:", currentDir)

	// Specify the file path (adjust as needed)
	filePath := "file.txt"

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist:", filePath)
		return
	}

	// Get the absolute path of the file
	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	fmt.Println("File exists. Absolute path:", absolutePath)
}
