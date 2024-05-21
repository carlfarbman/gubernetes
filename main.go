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

	// Convert the string into a slice of runes.
	runes := []rune(arg)

	// Check if the word has more than 2 characters.
	// If so, create an abbreviation by using the first and last character
	// with the rune count in between.
	// Otherwise, the abbreviation will be the word itself.
	var abbreviation string
	if len(runes) > 2 {
		abbreviation = fmt.Sprintf("%c%d%c", runes[0], len(runes)-2, runes[len(runes)-1])
	} else {
		abbreviation = arg
	}

	startBold := "\033[1m"
	// ANSI escape code to reset all formatting
	resetFormat := "\033[0m"

	// Print the bold abbreviation.
	// Wrap the abbreviation with the ANSI codes for bold text.
	fmt.Printf("%s: %s%s%s\n", arg, startBold, abbreviation, resetFormat)
}
