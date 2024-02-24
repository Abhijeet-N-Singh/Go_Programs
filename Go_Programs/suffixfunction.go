package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World!"
	suffix := "World!"
	if strings.HasSuffix(str, suffix) {
		fmt.Println("The string ends with given suffix.")
	} else {
		fmt.Println("The string does not end with given suffix.")
	}
}
