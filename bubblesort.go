package main

import "fmt"

// bubbleSort sorts a slice of integers in ascending order using the Bubble Sort algorithm.
func bubbleSort(slice []int) {
	n := len(slice)
	// Traverse through all array elements.
	for i := 0; i < n-1; i++ {
		// Last i elements are already in place.
		for j := 0; j < n-i-1; j++ {
			// Swap if the element found is greater than the next element.
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original slice:", numbers)

	bubbleSort(numbers)

	fmt.Println("Sorted slice:", numbers)
}
