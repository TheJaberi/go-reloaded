package reloaded

import (
	"strings"
)

func Capitalize(input string) string {
	capPattern := "(cap)"
	words := strings.Fields(input)
	for i := 1; i < len(words); i++ {
		if containsSubstring(words[i], capPattern) {
			words[i-1] = strings.Title(strings.ToLower(words[i-1]))
		}
	}
	marker := strings.Join(words, " ")
	converted := strings.ReplaceAll(marker, capPattern, "")
	nomarker := strings.Fields(converted)
	return strings.Join(nomarker, " ")
}
