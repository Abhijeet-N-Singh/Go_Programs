package main

// import the error package
import (
	"errors"
	"fmt"
)

func main() {

	message := "Hello Abhijeet Singh, How are you!"

	// create error using New() function
	myError := errors.New("The entered message is not matched with the actual message!!!!")

	if message != "Hell Ruhi Kumari, How are you!" {
		fmt.Println(myError)
	}
}
