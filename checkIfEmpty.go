package reloaded

import "strings"

func CheckIfEmpty(input string) (bool, int) {
	if len(input) == 0 {
		return true, 1
	} else if strings.TrimSpace(input) == "" {
		return true, 2
	} else {
		return false, 0
	}
}
