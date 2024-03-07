package main

import (
	"fmt"
)

func main() {
	var x = map[string]map[string]string{}

	x["fruits"] = map[string]string{}
	x["colors"] = map[string]string{}

	x["fruits"]["a"] = "apple"
	x["fruits"]["b"] = "banana"
	x["fruits"]["c"] = "cnataloupe"
	x["fruits"]["d"] = "dragon fruit"
	x["fruits"]["e"] = "elderberry"

	x["colors"]["r"] = "red"
	x["colors"]["g"] = "green"
	x["colors"]["b"] = "blue"
	x["colors"]["w"] = "white"
	x["colors"]["y"] = "yellow"

	fmt.Println(x)

}
