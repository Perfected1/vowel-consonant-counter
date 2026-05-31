package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Welcome message for the user
	fmt.Println("=== Vowel & Consonant Counter ===")

	// Create reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Prompt user
	fmt.Print("Enter a word or sentence: ")

	// Read full input line
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Clean input
	input = strings.TrimSpace(input)

	// Run counter logic
	vowels, consonants := countVowelsAndConsonants(input)

	// Output results
	fmt.Println("\nResults:")
	fmt.Println("Vowels:", vowels)
	fmt.Println("Consonants:", consonants)
}

// countVowelsAndConsonants processes the string and returns counts
func countVowelsAndConsonants(text string) (int, int) {
	vowelCount := 0
	consonantCount := 0

	// Define vowel set for quick lookup
	vowels := "aeiou"

	// Convert to lowercase for uniform comparison
	text = strings.ToLower(text)

	// Loop through each character
	for _, char := range text {
		// Ignore non-letter characters (spaces, numbers, symbols)
		if !unicode.IsLetter(char) {
			continue
		}

		// Check if character is a vowel
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else {
			consonantCount++
		}
	}

	return vowelCount, consonantCount
}