package reloaded

import (
	"strconv"
	"strings"
)

func BintoDec(Binstring string) string {
    var output int64
    words := strings.Fields(Binstring)
    for i := 1; i < len(words); i++ {
        if strings.ToLower(words[i]) == "(bin)" {
            number := strings.ToLower(words[i-1])
            words[i] = ""

            var err error
            output, err = strconv.ParseInt(number, 2, 64)
            if err != nil {
                // return an empty string in case of an error
                return ""
            }
            words[i-1] = strconv.FormatInt(output, 10)
        }
    }
    temp := strings.Join(words, " ")
    words = strings.Fields(temp)
    return  strings.Join(words, " ")
}