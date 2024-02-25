// Float bit size of 64:
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "3.142"
	result, err := strconv.ParseFloat(str, 64)

	if err == nil {
		fmt.Println("The float value is :", result)
	} else {
		fmt.Println("There is an error converting string to float.")
	}
}
