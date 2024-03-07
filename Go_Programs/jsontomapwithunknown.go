package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{
		"fruits" : {
			"a": "apple",
			"b": "banana",
			"c": "cantalouupe",
			"d": "dragon fruit",
			"e": " "elderberry"
		},
		"colors" : {
			"r": "red",
			"g": "green",
			"b": "blue",
			"w": "white",
			"y": "yellow"
		}
	}`
	var x map[string]interface{}

	json.Unmarshal([]byte(jsonStr), &x)
	fmt.Println(x)
}
