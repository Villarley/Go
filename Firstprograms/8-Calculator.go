package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	var res int
	var selectedOperation = askUserOp()
	a, b = obtainValues()
	// fmt.Println(a, b, selectedOperation)
	res, err := performOperation(selectedOperation, a, b)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The result is:", res)
	}
}
func askUserOp() int {
	var operation int
	var input string
	for {
		fmt.Println("Enter the operation you want to perform:\n1. Add\n2.Substract\n3.Multiply\n4.Divide\n")
		fmt.Scanln(&input)

		var err error
		operation, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
		} else {
			break // Break the loop
		}
	}
	return operation
}

func obtainValues() (int, int) {
	var a, b int
	var input string
	for {
		fmt.Println("Enter a number")
		fmt.Scanln(&input)

		// Convert the value into a number
		var err error
		a, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
		}
		fmt.Println("Enter a number")
		fmt.Scanln(&input)

		// Convert the value into a number
		b, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
		} else {
			break // Break the loop
		}
	}
	return a, b
}

func performOperation(selectedOperation int, a int, b int) (int, error) {
	var res int

	switch selectedOperation {
	case 1:
		res = Add(a, b)
	case 2:
		res = Substract(a, b)
	case 3:
		res = Multiply(a, b)
	case 4:
		if b == 0 {
			return 0, fmt.Errorf("Cannot divide by zero")
		}
		res = Divide(a, b)
	default:
		return 0, fmt.Errorf("Not valid")
	}
	return res, nil
}

func Add(a, b int) int {
	return a + b
}
func Substract(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) int {
	return a / b
}
