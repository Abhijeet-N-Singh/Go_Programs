package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	// Input from the user
	fmt.Println("Enter the value of num1: ")
	fmt.Scan(&num1)
	fmt.Println("Enter the value of num2: ")
	fmt.Scan(&num2)

	// Addition
	result1 := num1 + num2
	fmt.Printf("\nResult of num1 + num2 = %d\n", result1)

	// Substraction
	result2 := num1 - num2
	fmt.Printf("\nResult of num1 - num2 = %d\n", result2)

	// Multiplication
	result3 := num1 * num2
	fmt.Printf("\nResult of num1 * num2 = %d\n", result3)

	// Division
	result4 := num1 / num2
	fmt.Printf("\nResult of num1 / num2 = %d\n", result4)

	// Modulus
	result5 := num1 % num2
	fmt.Printf("\nResult of num1 %% num2 = %d\n", result5)

}
