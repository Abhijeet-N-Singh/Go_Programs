package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "Hello"
	var str2 = "World"
	var output = strings.Join([]string{str1, str2}, " ")
	fmt.Println(output)
}
