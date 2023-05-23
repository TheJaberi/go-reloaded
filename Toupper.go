package reloaded

import (
	"strings"
)

func ToUpper(input string) string {
	upPattern := "(up)"
	words := strings.Fields(input)
	for i := 1; i < len(words); i++ {
		if containsSubstring(words[i], upPattern) {
			words[i-1] = strings.ToUpper(words[i-1])
		}
	}
	marker := strings.Join(words, " ")
	converted := strings.ReplaceAll(marker, upPattern, "")
	nomarker := strings.Fields(converted)
	return strings.Join(nomarker, " ")
}
