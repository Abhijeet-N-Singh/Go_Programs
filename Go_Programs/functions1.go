package main

import "fmt"

func calculate(num1 int, num2 int) (int, int) {
	var add int
	var mul int
	add = num1 + num2
	mul = num1 * num2
	return add, mul
}

func main() {
	var add, mul = calculate(3, 8)
	fmt.Println("Addition is:", add, "\nMultiplication is", mul)
}
