package reloaded

import (
    "strings"
)

func FormatPunctuation(s string) string {
    x := strings.Fields(s)
    punctuations := ".,:;!?"
    for i, word := range x {
        if strings.ContainsAny(word, punctuations) {
            x[i-1] = x[i-1] + word
      x[i] = ""
        }
    }

  y := strings.Join(x, " ")
  z := strings.Fields(y)
    return strings.Join(z, " ")
}
