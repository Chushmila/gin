package main

import (
	"fmt"
	"math"
)

// findLargest returns the largest number in the given slice of integers.
func findLargest(nums []int) int {
	if len(nums) == 0 {
		// Return the smallest integer value if the slice is empty.
		return math.MinInt
	}

	// Initialize the largest number with the first element of the slice.
	largest := nums[0]

	// Iterate through the slice starting from the second element.
	for _, num := range nums[1:] {
		if num > largest {
			largest = num
		}
	}

	return largest
}

func main() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("The largest number is:", findLargest(numbers)) // Output: The largest number is: 9
}
