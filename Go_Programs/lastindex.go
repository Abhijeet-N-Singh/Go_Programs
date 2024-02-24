package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "Hello Hello Hello Hello Hello Hello"
	var substr = "Hello"
	var index = strings.LastIndex(str, substr)
	fmt.Println("The index of last occurrence of substring is: ", index)
}
