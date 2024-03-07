package main

import "fmt"

func main() {
	map1 := map[string]string{
		"a": "apple",
		"b": "ball",
		"c": "cat",
		"d": "dog",
		"e": "element",
	}

	var Key string
	fmt.Println("Enter the Key string to check it is present in Map or not?")
	fmt.Scan(&Key)

	_, isPresent := map1[Key]

	if isPresent {
		fmt.Print("Te Key is present in map.")
	} else {
		fmt.Print("The Key is not present in map. ")
	}
}
