package main

import "fmt"

func main() {

	//declaration and initialization
	var numbers = make([]int, 5, 10)

	fmt.Println("Size of Slice: ", len(numbers))
	fmt.Println("Capacity of Slice: ", cap(numbers))
}
