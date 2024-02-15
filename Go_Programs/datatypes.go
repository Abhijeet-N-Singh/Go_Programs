package main

import (
	"fmt"
	"reflect"
)

var print = fmt.Println

func main() {
	var x int = 65
	fmt.Println(reflect.TypeOf(x))
	fmt.Println("The value stored in x is:", x)
	var y float32 = 3.142
	fmt.Println(reflect.TypeOf(y))
	fmt.Println("The value stored in y is:", y)
	var z byte = 87
	fmt.Println(reflect.TypeOf(z))
	fmt.Println("The value stored in z is", z)
	var p rune = 365
	fmt.Println(reflect.TypeOf(p))
	fmt.Println("The value stored in p is:", p)
}
