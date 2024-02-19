package main

import "fmt"

func main() {
	var n int = 0
	var today int = 2

	switch today {
	case 1:
		fmt.Println("Today is Monday\n")
		fmt.Println("This statement is not related to switch case!")
	case 2:
		fmt.Println("Today is Tuesday\n")
		if n == 0 {

		}
		fmt.Println("This statement is not related to switch case!")
	case 3:
		fmt.Println("Today is Wednesday\n")
		fmt.Println("This statement is not related to switch case!")
	default:
		fmt.Println("This statement is not related to switch case!")
	}
}
