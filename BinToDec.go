package reloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BintoDec converts binary numbers enclosed in "(bin)" markers to decimal format.
func BintoDec(Binstring string) string {
	var output int64
	words := strings.Fields(Binstring) // Split the input string into individual words
	for i := 1; i < len(words); i++ {
		if containsSubstring(strings.ToLower(words[i]), "(bin)") {
			number := strings.ToLower(words[i-1]) // Extract the previous word as the binary number
			words[i] = ""                         // Remove the "(bin)" marker

			var err error
			output, err = strconv.ParseInt(number, 2, 64) // Convert binary number to decimal
			if err != nil {
				fmt.Println("You have a character that is invalid or a number that overflowed before the (bin) marker, check your input and try again")
				os.Exit(0) // Exit the program if there is an error during conversion
			}

			words[i-1] = strconv.FormatInt(output, 10) // Replace the binary number with its decimal representation
		}
	}
	temp := strings.Join(words, " ") // Join the modified words back into a string
	words = strings.Fields(temp)     // Split the modified string into words again to remove extra spaces
	return strings.Join(words, " ")  // Join the words back into a string
}
