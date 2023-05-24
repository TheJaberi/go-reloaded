package reloaded

import (
	"regexp"
)

func FormatPunctuation(str string) string {
	// Remove unwanted spaces before punctuations
	re := regexp.MustCompile(`\s*([.,:;])\s*`)
	str = re.ReplaceAllString(str, "$1 ")

	// Handle exclamations and questions
	reExclQuest := regexp.MustCompile(`\s*([!?])\s*`)
	str = reExclQuest.ReplaceAllString(str, "$1")

	// Handle ellipsis separately
	reEllipsis := regexp.MustCompile(`\s*\.\s*\.\s*\.\s*`)
	str = reEllipsis.ReplaceAllString(str, "... ")

	return str
}
