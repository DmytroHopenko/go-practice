package main

import (
	"fmt"
)

func calculate(num1 float64, num2 float64, operator string) float64 {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Can't devide by 0")
			return 0
		}
		return num1 / num2
	default:
		return 0
	}
}

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator: ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Printf("Sum: %.2f\n", calculate(num1, num2, operator))
}
