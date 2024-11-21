package main

import (
	"fmt"
	"math"
)

// sumOfDigits calculates the sum of the digits of a given integer.
func sumOfDigits(n int) int {
	// Use absolute value to handle negative numbers.
	n = int(math.Abs(float64(n)))

	sum := 0

	// Loop until n becomes 0.
	for n > 0 {
		// Add the last digit to sum.
		sum += n % 10
		// Remove the last digit from n.
		n /= 10
	}

	return sum
}

func main() {
	fmt.Println(sumOfDigits(1234))  // Output: 10
	fmt.Println(sumOfDigits(-5678)) // Output: 26
}
