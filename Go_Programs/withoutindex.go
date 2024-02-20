package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Element")
	for _, element := range x {
		fmt.Println(" ", element)
	}
}
