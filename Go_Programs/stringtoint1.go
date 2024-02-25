package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "14sixsixsixsix"
	// convert string to integer
	number, error := strconv.Atoi(str)
	fmt.Println("error:", error)
	fmt.Println("number:", number)
}
