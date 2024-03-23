package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	//print array elements
	fmt.Println("The created Array is: ", numbers)

	// Sum of array elements
	sum := 0
	for _, num := range numbers {
		sum += num

	}
	fmt.Printf("The sum of array elements is: %d\n", sum)

	// Array Length
	fmt.Printf("The Length of array is: %d\n", len(numbers))

	// Accessing elements by index
	fmt.Printf("Elements at index 0: %d\n", numbers[0])
	fmt.Printf("Elements at index 0: %d\n", numbers[1])
	fmt.Printf("Elements at index 0: %d\n", numbers[2])
	fmt.Printf("Elements at index 0: %d\n", numbers[3])
	fmt.Printf("Elements at index 0: %d\n", numbers[4])
}
