package main

import "fmt"

func main() {
	var n int = 5
	var i int = 0
	for i < n {
		var k int = 0
		for k < n {
			fmt.Println("* ")
			if i == k {
				break

			}
			k = k + 1
		}
		fmt.Println()
		i = i + 1
	}
}
