package reloaded

import (
	"strings"
)

// Capitalize capitalizes the word preceding the "(cap)" marker.
func Capitalize(input string) string {
	capPattern := "(cap)"          // The pattern to identify the "(cap)" marker
	words := strings.Fields(input) // Split the input string into individual words
	for i := 1; i < len(words); i++ {
		if containsSubstring(words[i], capPattern) {
			// Capitalize the preceding word
			words[i-1] = strings.Title(strings.ToLower(words[i-1]))
		}
	}
	marker := strings.Join(words, " ")                      // Join the modified words with markers back into a string
	converted := strings.ReplaceAll(marker, capPattern, "") // Remove the "(cap)" markers
	nomarker := strings.Fields(converted)                   // Split the modified string into words again
	return strings.Join(nomarker, " ")                      // Join the words back into a string
}
