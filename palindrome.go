package main

import (
	"fmt"
	"unicode"
)

// isPalindrome checks if the given string is a palindrome.
func isPalindrome(s string) bool {
	// Convert the string to lowercase and remove any non-letter characters.
	cleaned := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	// Initialize two pointers: one at the start and one at the end of the slice.
	left, right := 0, len(cleaned)-1

	// Compare characters from the start and end, moving towards the center.
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("radar")) // Output: true
	fmt.Println(isPalindrome("hello")) // Output: false
}
