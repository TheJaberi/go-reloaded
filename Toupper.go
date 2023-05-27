package reloaded

import (
	"strings"
)

// ToUpper converts the word preceding the "(up)" marker to uppercase.
func ToUpper(input string) string {
	upPattern := "(up)"            // The pattern to identify the "(up)" marker
	words := strings.Fields(input) // Split the input string into individual words
	for i := 1; i < len(words); i++ {
		if containsSubstring(words[i], upPattern) {
			// Convert the preceding word to uppercase
			words[i-1] = strings.ToUpper(words[i-1])
		}
	}
	marker := strings.Join(words, " ")                     // Join the modified words with markers back into a string
	converted := strings.ReplaceAll(marker, upPattern, "") // Remove the "(up)" markers
	nomarker := strings.Fields(converted)                  // Split the modified string into words again, for spacing
	return strings.Join(nomarker, " ")                     // Join the words back into a string
}
