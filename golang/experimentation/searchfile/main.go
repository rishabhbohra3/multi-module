package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func searchForFile(rootDir, targetFileName string) []string {
	filePathList := []string{}
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error walking the path: %v\n", err)
			return nil
		}

		if !info.IsDir() && info.Name() == targetFileName {
			fmt.Printf("Found the file: %s\n", path)
			filePathList = append(filePathList, path)
		}

		return nil
	})

	// fmt.Println(filePathList)

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}

	return filePathList
}

func manipulatePath(rootDir, command string) (string, string) {
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
	rootDir := "/Users/rishabhbohra/Documents/bito-code/MyProjects/GoLang/experimentation/data/" // Replace with the actual root directory path
	CommandName := "fb.pr_review"                                                                // Replace with the target file name
	rootDir, targetFileName := manipulatePath(rootDir, CommandName)
	fmt.Println("rootDir: "+rootDir, "fileName: "+targetFileName)
	filepathList := searchForFile(rootDir, targetFileName)
	fmt.Println("Final Output")
	fmt.Println(filepathList)
}
