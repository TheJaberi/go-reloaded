package reloaded

import (
	"regexp"
)

func FormatPunctuation(str string) string {
	// Remove unwanted spaces before punctuations
	space := regexp.MustCompile(`\s*([.,:;?!])`)
	str = space.ReplaceAllString(str, "$1")

	// Regular expression to match a punctuation excluding ellipsis, apostrophe and single quote, followed by a letter
	addspace := regexp.MustCompile(`([^'][.,:;?!])([[:alpha:]])`)

	// Replace matched patterns with the addition of a space
	str = addspace.ReplaceAllString(str, "$1 $2")

	return str
}
