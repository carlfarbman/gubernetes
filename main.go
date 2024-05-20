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

	// Convert the string into a slice of runes
	runes := []rune(arg)

	// Print the required output: word and the number of characters (as runes).
	fmt.Printf("Word: %s\nChars: %d\n", arg, len(runes))
}
