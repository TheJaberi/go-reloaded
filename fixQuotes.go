package reloaded

import "strings"

// NormalizeArabicQuotes converts Arabic single quotes to normal single quotes.
func NormalizeArabicQuotes(input string) string {
	// Define the Arabic single quotes and their corresponding replacements
	arabicQuotes := map[string]string{
		"‘": "'",
		"’": "'",
		"‚": "'",
		"‛": "'",
		"“": "\"",
		"”": "\"",
	}

	// Replace Arabic single quotes with normal single quotes
	for arabic, normal := range arabicQuotes {
		input = strings.ReplaceAll(input, arabic, normal)
	}

	return input
}
