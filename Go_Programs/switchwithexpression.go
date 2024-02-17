package main

import "fmt"

func main() {
	var today int = 5

	switch today {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Thrusday")
	case 4:
		fmt.Println("Today is Friday")
	case 5:
		fmt.Println("Today is Saturday")
	case 6:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Value for today is invalid")
	}
}
