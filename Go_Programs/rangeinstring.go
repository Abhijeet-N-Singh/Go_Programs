package main

import "fmt"

func main() {

	var str = "ABCDEFabcdef"

	// using index returned from range
	for index := range str {
		fmt.Printf("str[%d] = %d \n", index, str[index])
	}
}
