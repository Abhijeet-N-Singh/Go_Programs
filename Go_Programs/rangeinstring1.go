package main

import "fmt"

func main() {

	var str = "ABCDEFabcdef"

	// using index and element returned from range
	for index, element := range str {
		fmt.Printf("str[%d] = %d \n", index, element)
	}
}
