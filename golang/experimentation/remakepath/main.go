package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ManipulatePath(rootDir, command string) (string, string) {
	// Replace dots with path separators ("/") and append ".yaml" to the command
	parts := strings.Split(command, ".")
	for i, part := range parts {
		fmt.Println("part: " + part)
		parts[i] = filepath.Join(parts[i])
	}
	fileName := parts[len(parts)-1] + ".yaml"
	parts = parts[:len(parts)-1]

	// Join the parts and rootDir to construct the output directory
	outputDir := filepath.Join(rootDir, filepath.Join(parts...))

	return outputDir, fileName
}

func main() {
	rootDir := "/opt/automation"
	command := "pr_review"

	outputDir, fileName := ManipulatePath(rootDir, command)

	fmt.Printf("rootDir = %s and fileName = %s\n", outputDir, fileName)
}
