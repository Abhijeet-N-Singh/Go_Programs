package main

import "fmt"

func main() {
	i := 0
x:
	for ; i < 5; i++ {
		for k := 0; k < 5; k++ {
			if k == i {
				i++
				fmt.Println(" *")
				goto x
			}
			fmt.Print("    *")
		}
	}
}
