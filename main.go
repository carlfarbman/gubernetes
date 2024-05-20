package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there is exactly one argument provided.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <single-word-argument-here>")
		os.Exit(1)
	}

	// Grab the argument from the command line.
	arg := os.Args[1]

	// Print the required output: word and the number of characters.
	fmt.Printf("%s: number of chars in the string %d\n", arg, len(arg))
}
