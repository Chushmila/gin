package main

import "fmt"

// printMultiplicationTable prints the multiplication table (from 1 to 10) for the given number.
func printMultiplicationTable(n int) {
	fmt.Printf("Multiplication table for %d:\n", n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)
	printMultiplicationTable(number)
}
