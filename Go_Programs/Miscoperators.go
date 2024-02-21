package main

import "fmt"

func main() {
	p := 10

	// using address of operator(&) and
	// pointer indirection(*) operator.

	fmt.Println("\nThe address of the variable 'p':\n", &p)

	q := &p

	fmt.Println("\nThe Address of the variable 'p' is assigned to variable 'q':\n", q)
	fmt.Println("\nValue stored at 'p' is copied to the variable of 'q': ", *q)

	*q = 5

	fmt.Println("\nThe new value '5' is copied to the address of 'q': ", *q)

	fmt.Println("\nValue stored at 'q' is copied to the variable of 'p': ", p)
}
