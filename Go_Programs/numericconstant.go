package main

import "fmt"

var print = fmt.Println

func main() {
	const c1 int = 55
	print("The integer constant value stored in c1 is:", c1)
	const c2 float64 = 3.142
	print("The floating constant value stored in c2 is:", c2)
}
