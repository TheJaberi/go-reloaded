package reloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BintoDec(Binstring string) string {
	var output int64
	words := strings.Fields(Binstring)
	for i := 1; i < len(words); i++ {
		if containsSubstring(strings.ToLower(words[i]), "(bin)") {
			number := strings.ToLower(words[i-1])
			words[i] = ""

			var err error
			output, err = strconv.ParseInt(number, 2, 64)
			if err != nil {
				fmt.Println("You have a character that is invalid or a number that overflowed before the (bin) marker, check your input and try again")
				os.Exit(0)
			}
			words[i-1] = strconv.FormatInt(output, 10)
		}
	}
	temp := strings.Join(words, " ")
	words = strings.Fields(temp)
	return strings.Join(words, " ")
}
