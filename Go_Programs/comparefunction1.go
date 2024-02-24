package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Println("Enter the first string to compare: ")
	fmt.Scan(&a)
	var b string
	fmt.Println("Enter the second string to compare: ")
	fmt.Scan(&b)

	if a == b {
		fmt.Println("The two strings, a and b, are equal.")
	} else if a < b {
		fmt.Println("String a is less than string b.")
	} else if a > b {
		fmt.Println("String a is greater than string b.")
	}
}
