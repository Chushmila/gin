package main

import "fmt"

// reverseString reverses the given string and returns the result.
func reverseString(s string) string {
	// Convert the string to a slice of runes to handle multi-byte characters correctly.
	runes := []rune(s)

	// Initialize two pointers: one at the start and one at the end of the slice.
	left, right := 0, len(runes)-1

	// Swap the characters at the left and right pointers until they meet in the middle.
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	// Convert the slice of runes back to a string and return it.
	return string(runes)
}

func main() {
	input := "hello"
	fmt.Println("Reversed:", reverseString(input)) // Output: "olleh"
}
