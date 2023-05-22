package reloaded

import "strings"

func Atoan(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			nextWord := strings.ToLower(words[i+1])
			if strings.HasPrefix(nextWord, "a") || strings.HasPrefix(nextWord, "e") || strings.HasPrefix(nextWord, "i") || strings.HasPrefix(nextWord, "o") || strings.HasPrefix(nextWord, "u") || strings.HasPrefix(nextWord, "h") {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}