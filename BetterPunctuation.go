package reloaded

import (
	"fmt"
	"strings"
)

func BetterPunctuation(input string) string {
	parts := strings.Split(input, "'")
	quoteCount := 0

	for i := 1; i < len(parts)-1; i += 2 {
		field := strings.Split(parts[i], "\\s+")

		// Join words with a space and trim spaces
		parts[i] = "'" + strings.TrimSpace(strings.Join(field, " ")) + "'"
		quoteCount++
	}
	if quoteCount%2 == 1 {
		fmt.Println("you have an odd number of single quotes, last single quote removed")
	}

	return strings.Join(parts, "")
}
