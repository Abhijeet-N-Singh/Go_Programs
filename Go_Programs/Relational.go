package main

import "fmt"

func main() {
	var p int
	var q int
	fmt.Println("Enter the value of p: ")
	fmt.Scan(&p)
	fmt.Println("Enter the value of q: ")
	fmt.Scan(&q)

	// '=='(Equal To)
	result1 := p == q
	fmt.Println("\nIs p is Equal to q: ", result1)

	// '!='(Not Equal To)
	result2 := p != q
	fmt.Println("\nIs p is Not Equal to q? ", result2)

	//'<'(Less Than)
	result3 := p < q
	fmt.Println("\nIs p is Less Than q ", result3)

	//'>'(Greater Than)
	result4 := p > q
	fmt.Println("\nIs p is Greater Than q? ", result4)

	//'>=' (Greater Than Equal To)
	result5 := p >= q
	fmt.Println("\nIs p is Greater Than Equal To q? ", result5)

	//'<=' (Less Than Equal To)
	result6 := p <= q
	fmt.Println("\nIs is Less Than Equal To q? ", result6)
}
