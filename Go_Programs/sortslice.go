package main

import (
	"fmt"
	"sort"
)

func main() {
	strSlice := []string{"e", "c", "d", "a", "b"}
	fmt.Println("Before sorting :=", strSlice)
	sort.Strings(strSlice)
	fmt.Println("After sorting :", strSlice)
}
