package main

import "fmt"

func main() {

	defer fmt.Println("one")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
}
