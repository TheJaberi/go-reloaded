package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := readTextFromFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	modifiedText := processText(text)
	err = writeTextToFile(modifiedText, outputFile)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Println("Text processing complete.")
}

func readTextFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeTextToFile(text, filename string) error {
	return ioutil.WriteFile(filename, []byte(text), 0644)
}

func processText(text string) string {
	// Replace (hex) with decimal version of the preceding word
	text = replaceHexNumbers(text)

	// Replace (bin) with decimal version of the preceding word
	text = replaceBinaryNumbers(text)

	// Convert the preceding word to uppercase if (up) is found
	text = convertToUppercase(text)

	// Convert the preceding word to lowercase if (low) is found
	text = convertToLowercase(text)

	// Capitalize the preceding word if (cap) is found
	text = convertToCapitalized(text)

	// Format punctuations with preceding and following words
	text = formatPunctuations(text)

	// Format the ' marks with the corresponding words
	text = formatQuotationMarks(text)

	// Replace 'a' with 'an' if the next word starts with a vowel or 'h'
	text = replaceAToAn(text)

	// Remove the marker tags from the text
	text = removeMarkers(text)

	// Normalize spacing between words
	text = normalizeSpacing(text)

	return text
}

func replaceHexNumbers(text string) string {
	re := regexp.MustCompile(`(\w+) \(hex\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		hex := strings.TrimSuffix(match, " (hex)")
		decimal := convertHexToDecimal(hex)
		return strings.Replace(match, hex, decimal, 1)
	})
}

func replaceBinaryNumbers(text string) string {
	re := regexp.MustCompile(`(\w+) \(bin\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		binary := strings.TrimSuffix(match, " (bin)")
		decimal := convertBinaryToDecimal(binary)
		return strings.Replace(match, binary, decimal, 1)
	})
}

func convertToUppercase(text string) string {
	re := regexp.MustCompile(`(\w+) \(up(?:, (\d+))?\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		word := submatches[1]
		count := submatches[2]

		if count == "" {
			return strings.Replace(match, word, strings.ToUpper(word), 1)
		}

		num, err := strconv.Atoi(count)
		if err != nil {
			return match
		}

		words := strings.Fields(text)
		updatedWords := make([]string, len(words))
		copy(updatedWords, words)
		var modifiedCount int
		for i := 0; i < len(updatedWords); i++ {
			if updatedWords[i] == word {
				updatedWords[i] = strings.ToUpper(updatedWords[i])
				modifiedCount++
				if modifiedCount == num {
					break
				}
			}
		}

		return strings.Join(updatedWords, " ")
	})
}

func convertToLowercase(text string) string {
	re := regexp.MustCompile(`(\w+) \(low(?:, (\d+))?\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		word := submatches[1]
		count := submatches[2]

		if count == "" {
			return strings.Replace(match, word, strings.ToLower(word), 1)
		}

		num, err := strconv.Atoi(count)
		if err != nil {
			return match
		}

		words := strings.Fields(text)
		updatedWords := make([]string, len(words))
		copy(updatedWords, words)
		var modifiedCount int
		for i := 0; i < len(updatedWords); i++ {
			if updatedWords[i] == word {
				updatedWords[i] = strings.ToLower(updatedWords[i])
				modifiedCount++
				if modifiedCount == num {
					break
				}
			}
		}

		return strings.Join(updatedWords, " ")
	})
}

func convertToCapitalized(text string) string {
	re := regexp.MustCompile(`(\w+) \(cap(?:, (\d+))?\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		word := submatches[1]
		count := submatches[2]

		if count == "" {
			return strings.Replace(match, word, strings.Title(word), 1)
		}

		num, err := strconv.Atoi(count)
		if err != nil {
			return match
		}

		words := strings.Fields(text)
		updatedWords := make([]string, len(words))
		copy(updatedWords, words)
		var modifiedCount int
		for i := 0; i < len(updatedWords); i++ {
			if updatedWords[i] == word {
				updatedWords[i] = strings.Title(updatedWords[i])
				modifiedCount++
				if modifiedCount == num {
					break
				}
			}
		}

		return strings.Join(updatedWords, " ")
	})
}

func formatPunctuations(text string) string {
	re := regexp.MustCompile(`(\w+)(?:\s*([.,!?;:]))(\s+|$)`)
	return re.ReplaceAllString(text, "$1$2 ")
}

func formatQuotationMarks(text string) string {
	re := regexp.MustCompile(`(\w+)\s+'([^']+)'\s+(\w+)`)
	return re.ReplaceAllString(text, "$1 '$2' $3")
}

func replaceAToAn(text string) string {
	re := regexp.MustCompile(`(?i)\ba\s+([aeiouh])`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		return strings.Replace(match, "a ", "an ", 1)
	})
}

func removeMarkers(text string) string {
	markerRe := regexp.MustCompile(`\((\w+)(?:,\s*\d+)?\)`)
	return markerRe.ReplaceAllString(text, "")
}

func normalizeSpacing(text string) string {
	// Remove extra spaces before punctuation marks
	text = regexp.MustCompile(`\s+([.,!?;:])`).ReplaceAllString(text, "$1")

	// Add space after punctuation marks if not followed by whitespace
	text = regexp.MustCompile(`([.,!?;:])\S`).ReplaceAllString(text, "$1 ")

	return text
}

func convertHexToDecimal(hex string) string {
	// Assuming hex is a valid hexadecimal number
	decimal := fmt.Sprintf("%d", convertBaseToDecimal(hex, 16))
	return decimal
}

func convertBinaryToDecimal(binary string) string {
	// Assuming binary is a valid binary number
	decimal := fmt.Sprintf("%d", convertBaseToDecimal(binary, 2))
	return decimal
}

func convertBaseToDecimal(number string, base int) int64 {
	// Assuming number is a valid number in the specified base
	var result int64
	for _, digit := range number {
		value := int64(digit) - '0'
		result = result*int64(base) + value
	}
	return result
}
