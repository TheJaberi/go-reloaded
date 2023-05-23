package main

import (
	"fmt"
	"os"
	"reloaded"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}
	var isEmpty bool
	var cases int
	isEmpty, cases = reloaded.CheckIfEmpty(string(text))

	if isEmpty {
		if cases == 1 {
			fmt.Println("you have an empty string in the input file, write some words 0_o")
		} else if cases == 2 {
			fmt.Println("you just have whitespaces in your input, write some words 0_o")
		} else {
			fmt.Println("you should check if your input is valid")
		}
		os.Exit(0)
	}

	modifiedText := reloaded.Atoan(string(text))
	modifiedText = reloaded.HexatoDec(modifiedText)
	modifiedText = reloaded.BintoDec(modifiedText)
	modifiedText = reloaded.Capitalize(modifiedText)
	modifiedText = reloaded.ToUpper(modifiedText)
	modifiedText = reloaded.MarkerWithNumber(modifiedText)
	modifiedText = reloaded.FormatPunctuation(modifiedText)
	modifiedText = reloaded.BetterPunctuation(modifiedText)

	err = os.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Println("Text processing complete. " + string(modifiedText))
}
