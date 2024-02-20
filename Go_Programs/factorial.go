package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int
	fmt.Println("Enter the Number: ")
	/** For Negative numbers we cannot calculate factorial **/
	fmt.Scan(&n)
	result := factorial(n)
	fmt.Println("Factorial of", n, "is :", result)
}
