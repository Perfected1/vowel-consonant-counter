package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"vowel-consonant-counter/counter"
)

func main() {
	// CLI flag for direct input
	textFlag := flag.String("text", "", "Text to analyze directly")
	flag.Parse()

	fmt.Println("=== Vowel & Consonant Counter ===")

	// If user passed text via flag, process immediately
	if *textFlag != "" {
		processText(*textFlag)
		return
	}

	// Otherwise enter interactive mode
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter text (or type 'exit' to quit): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)

		// Exit condition
		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		processText(input)
	}
}

// processText handles counting and output display
func processText(text string) {
	vowels, consonants := counter.CountVowelsAndConsonants(text)

	fmt.Println("\nResults:")
	fmt.Println("Vowels:", vowels)
	fmt.Println("Consonants:", consonants)
}