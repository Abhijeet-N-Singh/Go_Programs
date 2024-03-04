package main

import "fmt"

func main() {
	x := [6]int{10, 20, 30, 40, 50, 60}
	for index, element := range x {
		fmt.Printf("x[%d] = %d\n", index, element)
	}
}
