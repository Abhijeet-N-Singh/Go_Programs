package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World!"
	prefix := "Hello"

	if strings.HasPrefix(str, prefix) {
		fmt.Println("The string starts with given prefix.")
	} else {
		fmt.Println("The string does not start with gievn prefix.")
	}
}
