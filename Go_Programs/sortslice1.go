package main

import (
	"fmt"
	"sort"
)

func main() {
	strSlice := []string{"e", "c", "d", "a", "b"}
	fmt.Println("Before sorting :", strSlice)
	sort.Sort(sort.Reverse(sort.StringSlice(strSlice)))
	fmt.Println("After sorting :", strSlice)

}
