package main

import (
	"fmt"
)

func main() {
	var x = []string{"apple", "banana", "cherry", "grapes", "watermelon"}

	for i := 0; i < len(x); i++ {
		fmt.Println("The fruits present in the Slice are: ", x[i])
	}
}
