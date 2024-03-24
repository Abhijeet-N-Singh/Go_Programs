package main

import "fmt"

func main() {
	// Creating a map
	person := map[string]interface{}{
		"name":   "Abhijeet Singh Thakur",
		"age":    37,
		"gender": "Male",
	}

	// Print map eelements
	fmt.Println("The created Map is ", person)
	fmt.Println()

	// Accessing elements by key
	fmt.Printf("The Nmae of the person is: %s\n", person["name"])
	fmt.Println()

	// Adding a new key-value pair
	person["city"] = "New York"

	// Deleting a key-value pair
	delete(person, "age")

	fmt.Println("The Updated Map is: ", person)
}
