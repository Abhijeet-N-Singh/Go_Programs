package main

import "fmt"

func main() {
	var str1 string
	fmt.Println("Enter the first string to compare: ")
	fmt.Scan(&str1)
	var str2 string
	fmt.Println("Enter the second string to comapre: ")
	fmt.Scan(&str2)

	if str1 == str2 {
		fmt.Println("The two strings are equal.")
	} else {
		fmt.Println("The two strings are equal")
	}
}
