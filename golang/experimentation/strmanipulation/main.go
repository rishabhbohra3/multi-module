package main

import (
	"fmt"
	"io"

	"github.com/valyala/fasttemplate"
)

func replacePlaceholders(templateStr string, placeholderValues map[string]string) (string, error) {
	t, err := fasttemplate.NewTemplate(templateStr, "{{", "}}")
	if err != nil {
		return "", err
	}

	s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		println("value: " + placeholderValues[tag])
		return w.Write([]byte(placeholderValues[tag]))
	})
	return s, err
}

func main() {
	// Input data
	templateStr := "Hi {{name}}, welcome home. {{something}}"
	placeholderValues := map[string]string{"name": "Rishabh"}

	// Replace placeholders in the template
	result, err := replacePlaceholders(templateStr, placeholderValues)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("result:", result)
}
