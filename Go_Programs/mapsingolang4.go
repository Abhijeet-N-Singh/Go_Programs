package main

import "fmt"

func main() {
	var colorMap = map[string]string{"white": "#FFFFFF", "black": "#000000", "red": "#FF0000"}

	//update

	colorMap["red"] = "#FF2222"
	fmt.Println(colorMap["red"])
}
