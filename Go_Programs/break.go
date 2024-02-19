package main

import "fmt"

func main() {
	var i int = 0
	for i < 50 {
		if i == 5 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("If the value i is equal to 5 for loop stops it's iteration as we used break statement!")
}
