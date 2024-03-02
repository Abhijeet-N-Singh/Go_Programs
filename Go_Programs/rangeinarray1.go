package main

import "fmt"

func main() {
	var evens = [5]int{2, 4, 6, 8, 10}

	// using index and element returned from range

	for index, element := range evens {
		fmt.Printf("evens[%d] = %d \n", index, element)
	}
}
