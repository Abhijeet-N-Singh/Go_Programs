package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := map[string]string{
		"a": "apple",
		"b": "ball",
		"c": "cat",
		"d": "dog",
		"e": "element",
	}

	fmt.Println("Before converting into JSON String the given Map is: \n", map1)
	jsonStr, err := json.Marshal(map1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("\nAfter converting Map into JSON String is: \n", string(jsonStr))
	}
}
