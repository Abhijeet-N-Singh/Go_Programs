package main

import "fmt"

func main() {
	var p int
	var q int
	fmt.Println("Enter the value of p: ")
	fmt.Scan(&p)
	fmt.Println("Enter the value of q: ")
	fmt.Scan(&q)

	// Addition
	result1 := p + q
	fmt.Printf("\nResult  of p + q = %d\n", result1)

	// Substraction
	result2 := p - q
	fmt.Printf("\nResult  of p - q = %d\n", result2)

	// Multiplication
	result3 := p * q
	fmt.Printf("\nResult  of p * q = %d\n", result3)

	// Modulus
	result4 := p % q
	fmt.Printf("\nResult  of p %% q = %d\n", result4)

	//Division
	result5 := p / q
	fmt.Printf("\nResult  of p / q = %d\n", result5)

}
