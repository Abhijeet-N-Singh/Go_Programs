package main

import "fmt"

func main() {
	var numbers = []int{10, 20, 30, 40, 50}

	// accessing elements of slice using index
	num2 := numbers[2]
	num3 := numbers[3]

	fmt.Println("numbers[2] : ", num2)
	fmt.Println("numbers[3] : ", num3)
}
