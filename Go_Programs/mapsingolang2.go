package main

import "fmt"

func main() {
	var colorMap = make(map[string]string)
	colorMap["white"] = "#FFFFFF"
	colorMap["black"] = "#0000000"
	colorMap["red"] = "#FF0000"
	colorMap["blue"] = "#0000FF"
	colorMap["green"] = "#00FF00"

	fmt.Println(colorMap)
}
