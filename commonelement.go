package main

import "fmt"

// findCommonElements returns a slice containing the common elements between two given slices of integers.
func findCommonElements(slice1, slice2 []int) []int {
	commonMap := make(map[int]bool) // A map to track elements in the first slice.
	commonSlice := []int{}          // A slice to store common elements.

	// Add elements from the first slice to the map.
	for _, num := range slice1 {
		commonMap[num] = true
	}

	// Check elements of the second slice against the map.
	for _, num := range slice2 {
		// If the element exists in the map, it is a common element.
		if commonMap[num] {
			commonSlice = append(commonSlice, num)
			delete(commonMap, num) // Remove to avoid duplicates in the result.
		}
	}

	return commonSlice
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	common := findCommonElements(slice1, slice2)
	fmt.Println("Common elements:", common) // Output: Common elements: [4 5]
}
