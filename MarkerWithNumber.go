package reloaded

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// MarkerWithNumber applies transformations based on markers containing a number.
func MarkerWithNumber(input string) string {
	words := strings.Fields(input) // Split the input string into individual words

	noSpace := regexp.MustCompile(`\((low|up|cap)\d+\)`)  // Regex pattern to match markers without spaces
	noComma := regexp.MustCompile(`\((low|up|cap),\d+\)`) // Regex pattern to match markers without commas
	noBoth := regexp.MustCompile(`\((low|up|cap)`)        // Regex pattern to match markers without both spaces and commas

	for i := 0; i < len(words); i++ {
		if noSpace.MatchString(words[i]) || noComma.MatchString(words[i]) ||
			noBoth.MatchString(words[i]) && !containsSubstring(words[i], ",") {
			// Check for invalid markers without proper spacing or commas
			fmt.Println("you have an Invalid marker. A marker should contain both a comma and a space.")
			os.Exit(0) // Exit the program if an invalid marker is found
		}

		if strings.HasPrefix(words[i], "(low,") || strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(words[i], "(cap,") {
			// Check for markers with number prefixes (e.g., "(low, 3)", "(up, 2)", "(cap, 1)")

			parts := strings.Split(words[i], ",")
			if len(parts) == 2 {
				number, err := strconv.Atoi(words[i+1][:len(words[i+1])-1]) // Extract the number from the marker and convert it to an integer
				if number > i {
					fmt.Println("your number inside the up/low/cap marker is larger than number of words before it")
					os.Exit(0) // Exit the program if the number inside the marker is greater than the current index
				} else if err == nil && i-number >= 0 {
					// Perform transformations based on the marker and the number

					if strings.HasPrefix(words[i], "(low,") {
						// Convert the preceding words to lowercase
						for j := i - number; j < i; j++ {
							words[j] = strings.ToLower(words[j])
						}
					} else if strings.HasPrefix(words[i], "(up,") {
						// Convert the preceding words to uppercase
						for j := i - number; j < i; j++ {
							words[j] = strings.ToUpper(words[j])
						}
					} else if strings.HasPrefix(words[i], "(cap,") {
						// Capitalize the preceding words
						for j := i - number; j < i; j++ {
							words[j] = strings.Title(words[j])
						}
					}

					words[i+1] = ""                           // Clear the marker's number
					words[i] = ""                             // Clear the marker itself
					words = append(words[:i], words[i+1:]...) // Remove the cleared elements from the words slice
					i -= number                               // Step back since we removed an element
					temp := strings.Join(words, " ")          // Join the modified words back into a string
					words = strings.Fields(temp)              // Split the modified string into words again, for spacing
				}
			}
		}
	}

	return strings.Join(words, " ") // Join the words back into a string
}
