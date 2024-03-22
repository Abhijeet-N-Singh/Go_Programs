package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func main() {
	var n int
	// Input from the user
	fmt.Println("Enter the value of n: ")
	fmt.Scan(&n)
	fmt.Printf("Fibonacci series up to %d terms: ", n)
	fibonacci(n)
}
