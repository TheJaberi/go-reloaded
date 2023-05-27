package reloaded

import "strings"

// CheckIfEmpty checks if the input string is empty or contains only whitespaces.
// It returns a boolean indicating if the input is empty or contains only whitespaces,
// and an integer code representing the specific case.

func CheckIfEmpty(input string) (bool, int) {
	if len(input) == 0 { // Check if the input string has zero length
		return true, 1 // Return true and case 1: empty string
	} else if strings.TrimSpace(input) == "" { // Check if the input string, after trimming leading/trailing spaces, is empty
		return true, 2 // Return true and case 2: whitespace-only string
	} else {
		return false, 0 // Return false and case 0: input is not empty or whitespace-only
	}
}
