package main

import (
	"fmt"
)

func main() {
	var iterations int

	fmt.Println("Enter count of numbers: ")
	fmt.Scanln(&iterations)

	var oddNumbers = []int{}
	var evenNumbers = []int{}

	for i := 1; i <= iterations; i++ {
		var num int

		fmt.Printf("Enter number %d/%d\n", i, iterations)
		fmt.Scanln(&num)

		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		} else {
			oddNumbers = append(oddNumbers, num)
		}
	}

	fmt.Printf("Odd numbers: %v\n", oddNumbers)
	fmt.Printf("Even numbers: %v\n", evenNumbers)
}
