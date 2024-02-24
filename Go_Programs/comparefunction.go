package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Println("Enter the first string to compare: ")
	fmt.Scan(&a)
	var b string
	fmt.Println("Enter the second string to compare: ")
	fmt.Scan(&b)

	result := strings.Compare(a, b)
	if result == 0 {
		fmt.Println("The two strings, a and b, are equal.")
	} else if result == -1 {
		fmt.Println("String a is less than string b.")
	} else if result == 1 {
		fmt.Println("String a is greater than string b.")
	}
}
