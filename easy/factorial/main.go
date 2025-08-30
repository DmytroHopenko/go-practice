package main

import (
	"fmt"
)

func FactorialRecursive(num int) int {
	if num <= 1 {
		return 1
	}

	return num * FactorialRecursive(num-1)
}

func FactorialLoop(num int) int {
	if num < 1 {
		return 1
	}

	var result int = 1

	for i := num; i > 0; i-- {
		result *= i
	}

	return result
}

func main() {
	var num int

	fmt.Println("Enter number for factorial")
	fmt.Scanln(&num)

	fmt.Printf("\nFactorial using recursive: %v", FactorialRecursive(num))
	fmt.Printf("\n Factorial using loop %v\n", FactorialLoop(num))
}
