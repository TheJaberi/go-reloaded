package reloaded

import (
	"strings"
)

func ToLower(input string) string {
	lowPattern := "(low)"
	words := strings.Fields(input)
	for i := 1; i < len(words); i++ {
		if containsSubstring(words[i], lowPattern) {
			words[i-1] = strings.ToLower(words[i-1])
		}
	}
	marker := strings.Join(words, " ")
	converted := strings.ReplaceAll(marker, lowPattern, "")
	nomarker := strings.Fields(converted)
	return strings.Join(nomarker, " ")
}
