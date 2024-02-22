package main

import "fmt"

func add(num1 int, num2 int) int {
	var sum int
	sum = num1 + num2
	return sum
}

func main() {
	var sum1 = add(5, 10)
	fmt.Println("\nThe result of sum1 is: ", sum1)

	var sum2 = add(10, 17)
	fmt.Println("\nThe result of sum2 is: ", sum2)
}
