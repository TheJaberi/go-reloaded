package reloaded

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func MarkerWithNumber(input string) string {
	if containsSubstring(input, "(up, ") || containsSubstring(input, "(low, ") || containsSubstring(input, "(cap, ") {

		words := strings.Fields(input)
		normal := regexp.MustCompile(`\((low|up|cap)\)`)
		withNum := regexp.MustCompile(`\((low|up|cap),\s+\d+\)`)
		for i := 0; i < len(words); i++ {
			if !normal.MatchString(words[i]) {
				goto here
			} else if !withNum.MatchString(words[i]) {
				fmt.Println("you have an Invalid marker. A marker should contain both a comma and a space.")
				os.Exit(0)
			}
		here:
			if strings.HasPrefix(words[i], "(low,") || strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(words[i], "(cap,") {
				parts := strings.Split(words[i], ",")
				if len(parts) == 2 {
					number, err := strconv.Atoi(words[i+1][:len(words[i+1])-1]) // remove ')' and convert to int
					if number > i {
						fmt.Println("your number inside the up/low/cap marker is too large")
						os.Exit(0)
					} else if err == nil && i-number >= 0 {
						if strings.HasPrefix(words[i], "(low,") {
							for j := i - number; j < i; j++ {
								words[j] = strings.ToLower(words[j])
							}
						} else if strings.HasPrefix(words[i], "(up,") {
							for j := i - number; j < i; j++ {
								words[j] = strings.ToUpper(words[j])
							}
						} else if strings.HasPrefix(words[i], "(cap,") {
							for j := i - number; j < i; j++ {
								words[j] = strings.Title(words[j])
							}
						}
						words[i+1] = ""
						words[i] = ""
						words = append(words[:i], words[i+1:]...)
						i -= number // step back since we removed an element
					}
				}
			}
		}
		temp := strings.Join(words, " ")
		temp2 := strings.Fields(temp)
		return strings.Join(temp2, " ")

	}
	return input
}

//let it return errors if cap at beginning or number with cap too big
