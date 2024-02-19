package main

import "fmt"

func main() {
	var i int = 0
	for i < 10 {
		if i == 5 {
			i = i + 1
			fmt.Println("Because of continue statement, for loop iteration is skipped here when i value is equal to 5")
			continue
		}
		fmt.Println("The value of i is: ", i)
		i = i + 1
	}

}
