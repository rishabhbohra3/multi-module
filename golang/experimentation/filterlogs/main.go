package main

import (
	"fmt"
	"regexp"
	"strings"
)

func filterStrings(componentIds []string, inputString string) []string {
	var result []string

	// Split inputString into sessionID and commandName
	inputParts := strings.Split(inputString, ":")
	inputSessionID := inputParts[0]
	inputCommandName := inputParts[1]

	// Iterate through componentIds
	for _, componentID := range componentIds {
		// Split componentID into sessionID and commandName
		componentParts := strings.Split(componentID, ":")
		if len(componentParts) != 2 {
			// Skip invalid componentIds without sessionID and commandName
			continue
		}

		componentSessionID := componentParts[0]
		componentCommandName := componentParts[1]

		// Check conditions for filtering
		if inputSessionID == componentSessionID && strings.Contains(componentCommandName, inputCommandName) {
			result = append(result, componentID)
		}
	}

	return result
}

func extractValues(inputString string) ([]string, error) {
	// Define a regular expression pattern
	pattern := `\[(.*?)\]`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(inputString, -1)
	fmt.Println("matches: ", matches)

	// Check if there are matches
	if len(matches) == 0 {
		return nil, fmt.Errorf("no values found inside square brackets")
	}

	// Extract values from the matches
	var values []string
	for _, match := range matches {
		fmt.Println("match: ", match)
		if len(match) >= 2 {
			values = append(values, match[1])
		}
	}

	return values, nil
}

func extractComponentId(componentInfo string) string {
	// component id regular expression pattern
	pattern := `\[(.*?)\]`

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(componentInfo)

	fmt.Println(matches)
	if len(matches) < 2 {
		fmt.Println("no value found inside square brackets")
		return ""
	}

	value := matches[1]

	return value
}

func main() {

	rawString := "BITO_SESSION_MGR [] : found matching runId: ZPD1co1ZTGyZ2xE06_5frw for action [] [rishabh]"

	componentId := extractComponentId(rawString)
	fmt.Println(componentId)
	// Example input data
	componentIds := []string{"bito", "session-id", "session-id:com.example.bito", "session-id2:com.example.bito", "com.example.bito", componentId}
	inputString := "session-id:example"

	// Perform filtering
	output := filterStrings(componentIds, inputString)

	// Display the result
	fmt.Println("Filtered output:", output)

	values, err := extractValues("BITO_SESSION_MGR [sd] : found matching runId: ZPD1co1ZTGyZ2xE06_5frw for action [] [rishabh]")
	if err == nil {
		fmt.Println("values: ", values)
	} else {
		fmt.Println(err.Error())
	}

}
