package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"vowel-consonant-counter/counter"
)

func main() {
	// App header
	fmt.Println("=== Vowel & Consonant Counter ===")

	// Input reader
	reader := bufio.NewReader(os.Stdin)

	// Prompt user
	fmt.Print("Enter a word or sentence: ")

	// Read input
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Clean input
	input = strings.TrimSpace(input)

	// Call counter package
	vowels, consonants := counter.CountVowelsAndConsonants(input)

	// Display results
	fmt.Println("\nResults:")
	fmt.Println("Vowels:", vowels)
	fmt.Println("Consonants:", consonants)
} 