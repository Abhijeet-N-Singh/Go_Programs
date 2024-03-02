package main

import "fmt"

func main() {
	var colorMap = make(map[string]string)
	colorMap["white"] = "#FFFFFF"
	colorMap["black"] = "#000000"
	colorMap["red"] = "#FF0000"
	colorMap["blue"] = "#0000FF"
	colorMap["green"] = "#00FF00"

	for color, hex := range colorMap {
		fmt.Println("Hex value of", color, "is", hex)
	}
}
