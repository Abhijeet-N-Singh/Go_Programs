package main

import "fmt"

func main() {
	map1 := map[string]string{
		"a": "apple",
		"b": "ball",
		"c": "cat",
		"d": "dog",
		"e": "elephant",
	}

	map1["a"] = "avocado"

	fmt.Print(map1)

}
