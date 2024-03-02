package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40, 50}
	var element int
	fmt.Println("Enter the element that you want to check: ")
	fmt.Scan(&element)

	var result bool = false
	for _, x := range slice1 {
		if x == element {
			result = true
			break
		}
	}

	if result {
		fmt.Println("Element is present in the slice.")

	} else {
		fmt.Println("Element is not present in the slice.")
	}
}
