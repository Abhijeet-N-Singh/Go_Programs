package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "1566666"
	// convert string to integer
	number, error := strconv.Atoi(str)
	fmt.Println("error:", error)
	fmt.Println("number:", number)
}
