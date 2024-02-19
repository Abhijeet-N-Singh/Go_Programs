package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Index Element")
	for index, element := range x {
		fmt.Println(" ", index, " - ", element)
	}
}
