package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{5, 3, 4, 6, 2, 1}
	fmt.Println("Before sorting :=", intSlice)
	sort.Ints(intSlice)
	fmt.Println("After sorting :", intSlice)
}
