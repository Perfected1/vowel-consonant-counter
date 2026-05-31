package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Welcome message for the user
	fmt.Println("=== Vowel & Consonant Counter ===")

	// Create a reader to capture user input from terminal
	reader := bufio.NewReader(os.Stdin)

	// Ask the user for input
	fmt.Print("Enter a word or sentence: ")

	// Read input until newline
	input, err := reader.ReadString('\n')
	if err != nil {
		// Handle unexpected input error
		fmt.Println("Error reading input:", err)
		return
	}

	// Clean the input by removing spaces and newline characters
	input = strings.TrimSpace(input)

	// Show what the user entered (confirmation step)
	fmt.Println("\nYou entered:", input)

}