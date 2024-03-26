package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	srK := Employee{"Sr.", "Carstens", 1}

	fmt.Println("Sr. Carstens", srK)

	john := Employee{
		firstName: "John",
		lastName:  "Jhon",
		id:        2,
	}

	fmt.Println("John", john)

	var jane Employee
	jane.firstName = "Jane"
	jane.lastName = "Janettersons"
	jane.id = 3

	fmt.Println("Jane", jane)
}
