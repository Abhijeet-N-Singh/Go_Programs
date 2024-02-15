package main

import (
	"fmt"
	"reflect"
)

var print = fmt.Println

func main() {
	var x bool = true
	print(reflect.TypeOf(x))
	print("The value stored in x is:", x)
	var y bool = false
	print(reflect.TypeOf(y))
	print("The value stored in y is:", y)
}
