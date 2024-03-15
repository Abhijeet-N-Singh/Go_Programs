package main

import "fmt"

func main() {

	// create an empty interface
	var a interface{}

	// store integer to an empty interface
	a = "Golang"

	// type assertion
	interfaceValue := a.(int)

	fmt.Println(interfaceValue)
}
