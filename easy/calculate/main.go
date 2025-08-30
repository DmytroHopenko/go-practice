package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Printf("Sum: %d\n", num1+num2)
}
