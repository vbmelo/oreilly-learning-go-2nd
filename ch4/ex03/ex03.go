package main

import "fmt"

func main() {
	var total int

	// How the book wants me to do it
	// for i := 0; i < 10; i++ {
	//  This should be a =, not :=
	// 	total := total + i
	// 	fmt.Println(total)
	// }

	// How I would do it
	// for i := range 10 {
	// 	total := total + i
	// 	fmt.Println(total)
	// }

	// Correct way
	for i := range 10 {
		total += i
		fmt.Println(total)
	}
}
