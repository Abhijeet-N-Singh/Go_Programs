package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = []string{"US", "Canada", "Europe", "Australia"}
	var sep = ","
	var output = strings.Join(a, sep)
	fmt.Println(output)
}
