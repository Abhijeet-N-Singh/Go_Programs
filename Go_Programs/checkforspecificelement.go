package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	var element int
	fmt.Println("Enter the element that you want to check: ")
	fmt.Scan(&element)
	var result bool = false
	for _, x := range arr {
		if x == element {
			result = true
			break
		}
	}
	if result {
		fmt.Println("Element is present in the array.")
	} else {
		fmt.Println("Element is not present in the array.")
	}
}
