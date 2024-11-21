package main

import (
	"fmt"
	"unicode"
)

// countVowels counts the number of vowels (a, e, i, o, u) in a given string, case-insensitively.
func countVowels(s string) int {
	count := 0

	// Iterate over each character in the string.
	for _, char := range s {
		// Convert character to lowercase for case-insensitive comparison.
		switch unicode.ToLower(char) {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}

	return count
}

func main() {
	input := "Hello, World!"
	fmt.Printf("The number of vowels in \"%s\" is: %d\n", input, countVowels(input)) // Output: 3
}
