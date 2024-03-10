package main

import "fmt"

// define an error struct
type DivisonByZero struct {
	message string
}

// Create a function Error() string and associate it to the struct
func (z DivisonByZero) Error() string {
	return "Number Cannot Be Divided by Zero"
}

func divide(n1 int, n2 int) (int, error) {

	if n2 == 0 {
		return 0, &DivisonByZero{}
	} else {
		return n1 / n2, nil
	}
}

func main() {

	number1 := 15
	// change the value of number to get different result
	number2 := 0

	result, err := divide(number1, number2)

	// check if error occur or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d", result)
	}
}
