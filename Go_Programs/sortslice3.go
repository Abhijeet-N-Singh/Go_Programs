package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{5, 3, 4, 1, 2}
	fmt.Println("Before sorting :", intSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println("After sorting :", intSlice)

}
