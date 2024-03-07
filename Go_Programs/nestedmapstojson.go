package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x = map[string]map[string]string{}

	x["fruits"] = map[string]string{}
	x["colors"] = map[string]string{}

	x["fruits"]["a"] = "apple"
	x["fruits"]["b"] = "banana"
	x["fruits"]["c"] = "cantalope"
	x["fruits"]["d"] = "dragon fruit"
	x["fruits"]["e"] = "elderberry"

	x["colors"]["r"] = "red"
	x["colors"]["g"] = "green"
	x["colors"]["b"] = "blue"
	x["colors"]["w"] = "white"
	x["colors"]["y"] = "yellow"

	fmt.Println("Before converting into JSON String the given Nested Map is: \n", x)

	jsonStr, err := json.Marshal(x)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("\nAfter converting Nested Map into JSON String is: \n", string(jsonStr))
	}
}
