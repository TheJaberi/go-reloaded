package main

import (
	"fmt"
	"os"
	"reloaded"
)

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]  // Input file path
	outputFile := os.Args[2] // Output file path

	// Read the content of the input file
	text, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	var isEmpty bool
	var cases int

	// Check if the input text is empty or contains only whitespaces
	isEmpty, cases = reloaded.CheckIfEmpty(string(text))

	if isEmpty {
		// Handle the case of an empty or whitespace-only input
		if cases == 1 {
			fmt.Println("you have an empty string in the input file, write some words 0_o")
		} else if cases == 2 {
			fmt.Println("you just have whitespaces in your input, write some words 0_o")
		}
		os.Exit(0)
	}

	// Perform various modifications on the input text
	modifiedText := reloaded.Atoan(string(text))
	modifiedText = reloaded.HexatoDec(modifiedText)
	modifiedText = reloaded.BintoDec(modifiedText)
	modifiedText = reloaded.Capitalize(modifiedText)
	modifiedText = reloaded.ToUpper(modifiedText)
	modifiedText = reloaded.ToLower(modifiedText)
	modifiedText = reloaded.MarkerWithNumber(modifiedText)
	modifiedText = reloaded.FormatPunctuation(modifiedText)
	modifiedText = reloaded.BetterPunctuation(modifiedText)

	// Write the modified text to the output file
	err = os.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}
}
