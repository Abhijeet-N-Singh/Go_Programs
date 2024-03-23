package main

import "fmt"

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var input string
	fmt.Println("Enter the input string to reverse it.")
	fmt.Scan(&input)
	reversed := reverseString(input)
	fmt.Printf("Original: %s\nRversed: %s\n", input, reversed)
}
