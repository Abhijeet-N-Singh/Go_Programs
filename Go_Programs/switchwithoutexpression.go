package main

import "fmt"

func main() {
	var today int = 5

	switch {
	case today == 1:
		fmt.Println("Today is Monday")
	case today == 2:
		fmt.Println("Today is Tuesday")
	case today == 3:
		fmt.Println("Today is Thrusday")
	case today == 4:
		fmt.Println("Today is Friday")
	case today == 5:
		fmt.Println("Today is Saturday")
	case today == 6:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Value for today is invalid")
	}
}
