package reloaded

import (
	"regexp"
	"strings"
)

func BetterPunctuation(input string) string {
	re := regexp.MustCompile(`'([^']*)'`)
	matches := re.FindAllStringSubmatch(input, -1)

	quoteCount := 0
	for _, match := range matches {
		quoted := match[0]
		insideQuotes := match[1]

		field := strings.Split(insideQuotes, " ")

		// Join words with a space and trim spaces
		newQuoted := "'" + strings.TrimSpace(strings.Join(field, " ")) + "'"
		input = strings.Replace(input, quoted, newQuoted, 1)
		quoteCount++
	}

	return input
}
