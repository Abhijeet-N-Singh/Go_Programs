package main

// import the errors package
import (
	"errors"
	"fmt"
)

// function that checks if name is Ruhi Kumari
func checkName(name string) error {
	// create a new error
	newError := errors.New("Invalid Name")

	// return the error if name is not equal to Ruhi Kumar
	if name != "Ruhi Kumari" {
		return newError
	}

	// return nil if there is no error
	return nil
}

func main() {

	name := "Abhijeet Singh"

	// call the function
	err := checkName(name)

	// check if the err is nil or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid Name")
	}
}
