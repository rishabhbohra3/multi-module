package main

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID
	newUUID := uuid.New()

	// Convert UUID to a base64-encoded string
	base64UUID := base64.RawURLEncoding.EncodeToString(newUUID[:])

	// Print the generated UUID
	fmt.Println("Generated Base64 UUID:", base64UUID)

}
