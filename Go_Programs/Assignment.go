package main

import "fmt"

func main() {
	var p int
	var q int
	fmt.Println("Enter the value of p: ")
	fmt.Scan(&p)
	fmt.Println("Enter the value of q: ")
	fmt.Scan(&q)

	// " = " (single assignment)
	p = q
	fmt.Println("\nThe value of 'q' is assigned to 'p': ", p)

	// " += " (Add assignment)
	p += q
	fmt.Println("\nThe value of 'q' is added to 'p': ", p)

	// " -= " (Substract assignment)
	p -= q
	fmt.Println("\nThe value of 'q' is substracted from 'p': ", p)

	// " *= " (Multiply assignment)
	p *= q
	fmt.Println("\nThe value of 'q' is multiplied by 'p': ", p)

	// " /= " (Division assignment)
	p /= q
	fmt.Println("\nThe value of 'q' is divided by 'p': ", p)

	// " %= " (Modulus assignment)
	p %= q
	fmt.Println("\nThe value of 'q' is modulus of 'p': ", p)
}
