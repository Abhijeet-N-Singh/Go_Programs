// Negative Scenario:

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "3.onefourtwo"
	result, err := strconv.ParseFloat(str, 64)

	if err == nil {
		fmt.Println("The float value is :", result)
	} else {
		fmt.Println("There is an error cconverting string to float.")
		fmt.Println(err)
	}
}
