package main

import "fmt"

func main() {
	a := [6]int{52, 2, 14, 45, 6, 8}
	for i, x := range a {
		fmt.Println("The value at a[", i, "] is ", x, "\n")
	}
}
