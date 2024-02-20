package main

import (
	"fmt"
	"math/rand"
)

var names = []string{
	"Alpha", "Beta", "Gamma", "Delta", "Epsilon",
	"Zeta", "Eta", "Theta", "Iota", "Kappa",
	"Lambda", "Mu", "Nu", "Xi", "Omicron",
	"Pi", "Rho", "Sigma", "Tau", "Upsilon",
	"Phi", "Chi", "Psi", "Omega",
}

func getRandomName() string {
	// Include a formatted numerical part (e.g., 001, 012, 123)
	randomNumber := fmt.Sprintf("%03d", rand.Intn(10))

	// Select a random name from the list
	randomName := names[rand.Intn(len(names))]

	// Combine the name and formatted numerical part
	return fmt.Sprintf("%s%s", randomName, randomNumber)
}

func main() {
	// Seed is not required in Go 1.17 and later versions

	// Generate and print a random name
	randomName := getRandomName()
	fmt.Println("Random Name:", randomName)
}
