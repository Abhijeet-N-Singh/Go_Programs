package main

import "fmt"

//  struct definition
type Student struct {
	name       string
	rollNumber int
}

// associate method to Student struct type
func (s Student) PrintDetails() {
	fmt.Println("Student Details\n---------------")
	fmt.Println("Name :", s.name)
	fmt.Println("Roll Number :", s.rollNumber)
	fmt.Println("\n")
}

func main() {
	var student1 = Student{name: "Abhijeet Singh Thakur", rollNumber: 1206}
	student1.PrintDetails()

	var student2 = Student{name: "Ruhi Kumari", rollNumber: 1209}
	student2.PrintDetails()
}
