package main

import "fmt"

func main() {
	x := map[string]string{
		"a": "apple",
		"b": "ball",
		"c": "cat",
		"d": "dog",
		"e": "elephant",
	}

	for key := range x {
		fmt.Println(key)
	}
}
