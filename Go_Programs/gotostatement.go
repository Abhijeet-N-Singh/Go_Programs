package main

import (
	"fmt"
)

func main() {
	fmt.Println("Apple")
	goto x
	fmt.Println("Banana")
x:
	fmt.Println("Cheeeku")
}
