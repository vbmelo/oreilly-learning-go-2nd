package main

import "fmt"

func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	sub1 := greetings[:2]
	sub2 := greetings[1:4]
	sub3 := greetings[3:]

	fmt.Printf("sub1: %v\n sub2: %v\n sub3: %v\n", sub1, sub2, sub3)
}
