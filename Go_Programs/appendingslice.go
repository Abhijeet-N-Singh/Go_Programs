package main

import "fmt"

func main() {
	var numbers = []int{10, 20, 30, 40, 50}

	numbers = append(numbers, 60)

	fmt.Println(numbers)
}
