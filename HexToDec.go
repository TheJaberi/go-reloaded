package reloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HexatoDec converts hexadecimal numbers enclosed in "(hex)" markers to decimal format.
func HexatoDec(HextoDec string) string {
	var output int64
	words := strings.Fields(HextoDec) // Split the input string into individual words
	for i := 1; i < len(words); i++ {
		if containsSubstring(strings.ToLower(words[i]), "(hex)") {
			number := strings.ToLower(words[i-1]) // Extract the previous word as the hexadecimal number
			words[i] = ""                         // Remove the "(hex)" marker

			var err error
			output, err = strconv.ParseInt(number, 16, 64) // Convert hexadecimal number to decimal
			if err != nil {
				fmt.Println("You have a character that is invalid or a number that overflowed before the (hex) marker, check your input and try again")
				os.Exit(0) // Exit the program if there is an error during conversion
			}

			words[i-1] = strconv.FormatInt(output, 10) // Replace the hexadecimal number with its decimal representation
		}
	}
	temp := strings.Join(words, " ") // Join the modified words back into a string
	words = strings.Fields(temp)     // Split the modified string into words again
	return strings.Join(words, " ")  // Join the words back into a string
}
