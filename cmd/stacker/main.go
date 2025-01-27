package main

import (
	"fmt"
	"os"

	stackdecode "github.com/tmtf-stacker/stacker/internal/pkg/stack_decode"
)

func main() {
	// Check if the "configPath" argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <configPath>")
		os.Exit(1)
	}

	// Get the "configPath" argument
	configPath := os.Args[1]

	stackdecode.StackDecode(configPath)
}
