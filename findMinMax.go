package main

import (
	"fmt"
)

// findMinMax returns the minimum and maximum values from a given slice of integers.
func findMinMax(nums []int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, fmt.Errorf("slice is empty")
	}

	// Initialize min and max with the first element of the slice.
	min, max := nums[0], nums[0]

	// Iterate through the slice to find the min and max values.
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max, nil
}

func main() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	min, max, err := findMinMax(numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The minimum value is: %d\n", min) // Output: The minimum value is: 1
	fmt.Printf("The maximum value is: %d\n", max) // Output: The maximum value is: 9
}
