package main

import (
	"fmt"
	"reflect"
)

var print = fmt.Println

func main() {
	var x string = "Abhijeet Singh"
	print(reflect.TypeOf(x))
	print("The String stored in x is:", x)
	var y string = "Deepnisha Singh"
	print(reflect.TypeOf(y))
	print("The String stored in y is:", y)
}
