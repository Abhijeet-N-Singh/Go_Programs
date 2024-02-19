package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Index Element")
	for index, element := range x {
		fmt.Println(" ", index, " - ", element)
	}
}
