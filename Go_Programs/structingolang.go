package main

import "fmt"

type student struct {
	name       string
	rollNumber int
}

func main() {

	var student1 = student{name: "Abhijeet Singh", rollNumber: 1208}

	//get the member 'value'
	fmt.Printf("student1 details\n----------\n")
	fmt.Printf("Name : %s\n", student1.name)
	fmt.Printf("Roll Number : %d\n", student1.rollNumber)

	var student2 = student{name: "Ruhi Kumari", rollNumber: 1209}

	//before setting the number value
	fmt.Printf("\nstudent2 details\n---------------\n")
	fmt.Printf("Name : %s\n", student2.name)
	fmt.Printf("Roll Number : %d\n", student2.rollNumber)

	// after setting the number value
	fmt.Printf("\nstudent2 details\n-------------------\n")
	fmt.Printf("Name : %s\n", student2.name)
	fmt.Printf("Roll Number : %d\n", student2.rollNumber)
}
