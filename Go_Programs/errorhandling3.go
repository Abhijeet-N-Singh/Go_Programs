package main

import "fmt"

func main() {

	age := -14

	// create an error using Errorf()
	error := fmt.Errorf("%d is negative\n Age can't be negative", age)

	if age < 0 {
		fmt.Println(error)

	} else {
		fmt.Println("Age: %d", age)
	}

}
