package main

import "fmt"

func main() {
	var p int
	var q int
	fmt.Println("Enter the value of p: ")
	fmt.Scan(&p)
	fmt.Println("Enter the value of q: ")
	fmt.Scan(&q)

	// LOgical AND '&&'
	if p != q && p <= q {
		fmt.Println("Yes, it is true")
	} else {
		fmt.Println("NO, it is false")
	}

	// Logical OR '||'
	if p != q || p <= q {
		fmt.Println("Yes, it is true")
	} else {
		fmt.Println("NO, it is false")
	}

	// Logical NOT '!'
	if !(p == q) {
		fmt.Println("Yes , it is true")
	} else {
		fmt.Println("No,  it is false")
	}
}
