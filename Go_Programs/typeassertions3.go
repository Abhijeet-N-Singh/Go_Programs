package main

import "fmt"

func main() {

	// create an empty interface
	var a interface{}
	a = "Golang"

	// type assertion
	interfaceValue, booleanValue := a.(int)

	fmt.Println("Interface Value:", interfaceValue)
	fmt.Println("Boolean Value:", booleanValue)
}
