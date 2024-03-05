package main

import "fmt"

func main() {
	var colorMap = map[string]string{"white": "#FFFFFF", "black": "#0000000", "red": "#FF0000",
		"blue": "#0000ff", "green": "#00FF00"}

	delete(colorMap, "white")
	delete(colorMap, "black")

	fmt.Println(colorMap)
}
