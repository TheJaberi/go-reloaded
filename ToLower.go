package reloaded

import (
	"strings"
)

// ToLower converts the word preceding the "(low)" marker to lowercase.
func ToLower(input string) string {
	lowPattern := "(low)"          // The pattern to identify the "(low)" marker
	words := strings.Fields(input) // Split the input string into individual words
	for i := 1; i < len(words); i++ {
		if containsSubstring(words[i], lowPattern) {
			// Convert the preceding word to lowercase
			words[i-1] = strings.ToLower(words[i-1])
		}
	}
	marker := strings.Join(words, " ")                      // Join the modified words with markers back into a string
	converted := strings.ReplaceAll(marker, lowPattern, "") // Remove the "(low)" markers
	nomarker := strings.Fields(converted)                   // Split the modified string into words again, for spacing
	return strings.Join(nomarker, " ")                      // Join the words back into a string
}
