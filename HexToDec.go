package reloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HexatoDec(HextoDec string) string {
	var output int64
	words := strings.Fields(HextoDec)
	for i := 1; i < len(words); i++ {
		if containsSubstring(strings.ToLower(words[i]), "(hex)") {
			number := strings.ToLower(words[i-1])
			words[i] = ""

			var err error
			output, err = strconv.ParseInt(number, 16, 64)
			if err != nil {
				fmt.Println("You have a character that is invalid or a number that overflowed before the (hex) marker, check your input and try again")
				os.Exit(0)
			}
			words[i-1] = strconv.FormatInt(output, 10)
		}
	}
	temp := strings.Join(words, " ")
	words = strings.Fields(temp)
	return strings.Join(words, " ")
}
