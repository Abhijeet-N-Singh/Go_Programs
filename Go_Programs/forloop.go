package main

import "fmt"

func main() {
	var i int = 0
	for i < 5 {
		fmt.Println("The value of i is:", i)
		i = i + 1
	}
}
