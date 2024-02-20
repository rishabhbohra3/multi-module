package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Command to execute (adjust as needed)
	cmd := exec.Command("bito", "-p", "facts about golang")

	// pipe the commands output to the applications
	// standard output
	cmd.Stdout = os.Stdout

	// Run still runs the command and waits for completion
	// but the output is instantly piped to Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}
