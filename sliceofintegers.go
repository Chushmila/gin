package main

import "fmt"

// uniqueElements returns a new slice containing only the unique elements from the given slice of integers.
func uniqueElements(nums []int) []int {
	uniqueMap := make(map[int]bool) // A map to track unique elements.
	uniqueSlice := []int{}          // A slice to store unique elements.

	// Iterate through the given slice.
	for _, num := range nums {
		// Check if the element is already in the map.
		if !uniqueMap[num] {
			uniqueMap[num] = true                  // Mark the element as seen.
			uniqueSlice = append(uniqueSlice, num) // Add the element to the unique slice.
		}
	}

	return uniqueSlice
}

func main() {
	numbers := []int{1, 2, 3, 4, 2, 3, 5, 6, 7, 8, 1}
	unique := uniqueElements(numbers)
	fmt.Println("Unique elements:", unique) // Output: Unique elements: [1 2 3 4 5 6 7 8]
}
