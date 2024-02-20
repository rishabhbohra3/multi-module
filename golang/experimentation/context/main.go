package main

import (
	"context"
	"fmt"
	"reflect"
)

// ContextKey is a custom type for keys in the context
type ContextKey string

const (
	// KeyLoggerID is an example key for logger ID
	KeyLoggerID  ContextKey = "LoggerID"
	KeyLoggerID2 ContextKey = "LoggerID2"
)

func GetLoggerID(ctx context.Context) (string, error) {
	// Retrieve values from the context
	value := ctx.Value(KeyLoggerID)

	//Check if the value is not present in the context
	if value == nil {
		return "", fmt.Errorf("LoggerID not found in context")
	}

	// Check if the value is of type string
	if _, ok := value.(string); !ok {
		return "", fmt.Errorf("LoggerID in context has an unexpected type: %v", reflect.TypeOf(value))
	}

	// Convert the value to string
	loggerID := value.(string)

	// Check if the loggerID is not an empty string
	if loggerID == "" {
		return "", fmt.Errorf("LoggerID in context is empty")
	}

	// Return the validated logger ID
	return loggerID, nil
}

func main() {
	// Create a new context
	// ctx := context.WithValue(context.Background(), KeyLoggerID, 123) // Setting a non-string value

	ctx := context.WithValue(context.Background(), KeyLoggerID2, 123) // Setting a non-string value

	// Example usage of GetLoggerID
	loggerID, err := GetLoggerID(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the validated logger ID
	fmt.Println("LoggerID:", loggerID)
}
