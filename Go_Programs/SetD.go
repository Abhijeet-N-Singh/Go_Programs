package main

import "fmt"

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	var num int

	// Input form the user
	fmt.Println("Enter the value of num : ")
	fmt.Scan(&num)

	if num == 0 {
		fmt.Printf("%d is neither prime nor composite\n", num)
	} else if isPrime(num) {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime numbner\n", num)
	}
}
