package counter

import (
	"strings"
	"unicode"
)

// CountVowelsAndConsonants takes a string and returns vowel and consonant counts
func CountVowelsAndConsonants(text string) (int, int) {
	vowelCount := 0
	consonantCount := 0

	// Define vowels for comparison
	vowels := "aeiou"

	// Normalize input
	text = strings.ToLower(text)

	// Loop through each character
	for _, char := range text {

		// Skip anything that is not a letter
		if !unicode.IsLetter(char) {
			continue
		}

		// Check if vowel
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else {
			consonantCount++
		}
	}

	return vowelCount, consonantCount
}