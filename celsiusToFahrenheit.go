package main

import (
	"fmt"
)

// celsiusToFahrenheit converts Celsius to Fahrenheit.
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// fahrenheitToCelsius converts Fahrenheit to Celsius.
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var choice string
	var temperature float64

	// Loop until a valid choice is made.
	for {
		fmt.Println("Choose the conversion:")
		fmt.Println("1. Celsius to Fahrenheit")
		fmt.Println("2. Fahrenheit to Celsius")
		fmt.Print("Enter your choice (1 or 2): ")
		fmt.Scanf("%s", &choice)

		if choice == "1" || choice == "2" {
			break // Exit loop if the choice is valid.
		}
		fmt.Println("Invalid choice, please enter 1 or 2.")
	}

	fmt.Print("Enter the temperature: ")
	fmt.Scanf("%f", &temperature)

	if choice == "1" {
		// Convert Celsius to Fahrenheit.
		result := celsiusToFahrenheit(temperature)
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temperature, result)
	} else {
		// Convert Fahrenheit to Celsius.
		result := fahrenheitToCelsius(temperature)
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temperature, result)
	}
}
