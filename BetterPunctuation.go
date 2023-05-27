package reloaded

import (
	"regexp"
	"strings"
)

// BetterPunctuation improves the punctuation by modifying the quoted text within single quotes.
func BetterPunctuation(input string) string {
	re := regexp.MustCompile(`'([^']*)'`)          // Regular expression to match single quotes and the content within them
	matches := re.FindAllStringSubmatch(input, -1) // Find all matches of the regular expression in the input string

	for _, match := range matches {
		quoted := match[0]       // The matched portion within single quotes, including the quotes
		insideQuotes := match[1] // The content within the single quotes

		field := strings.Fields(insideQuotes) // Split the content into individual words

		// Join the words with a space and trim leading/trailing spaces
		newQuoted := "'" + strings.Join(field, " ") + "'" // Create the modified quoted text with improved punctuation

		// Replace the original quoted text with the modified one in the input string
		input = strings.Replace(input, quoted, newQuoted, 1)
	}

	return input
}
