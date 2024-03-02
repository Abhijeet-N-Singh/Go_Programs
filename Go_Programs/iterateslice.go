package main

import (
	"fmt"
)

func main() {
	var x = []int{10, 20, 30, 40, 50}

	for i := 0; i < len(x); i++ {
		fmt.Println("The elements present in Slice are: ", x[i])
	}

}
