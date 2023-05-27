package reloaded

import "strings"

// Atoan converts instances of the word "a" to "an" if the following word starts with a vowel or the letter "h".
func Atoan(input string) string {
	words := strings.Fields(input) // Split the input into individual words
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" { // Check if the current word is "a"
			nextWord := strings.ToLower(words[i+1]) // Get the next word and convert it to lowercase
			if strings.HasPrefix(nextWord, "a") || strings.HasPrefix(nextWord, "e") ||
				strings.HasPrefix(nextWord, "i") || strings.HasPrefix(nextWord, "o") ||
				strings.HasPrefix(nextWord, "u") || strings.HasPrefix(nextWord, "h") {
				words[i] = "an" // Replace "a" with "an"
			}
		}
	}
	return strings.Join(words, " ") // Join the modified words back into a string, then return
}
