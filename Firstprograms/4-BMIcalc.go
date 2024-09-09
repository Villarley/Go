package main

import (
	"fmt"
	"strconv"
)

func main() {
	var heightStr, weightStr string
	fmt.Print("Enter your height in meters: ")
	fmt.Scanln(&heightStr)
	fmt.Print("Enter your weight in kilograms: ")
	fmt.Scanln(&weightStr)

	// Convert string to float64
	height, err1 := strconv.ParseFloat(heightStr, 64)
	weight, err2 := strconv.ParseFloat(weightStr, 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid input.")
		return
	}

	BMI := calculateBMI(height, weight)
	fmt.Printf("Your BMI is: %f\n", BMI)
}

func calculateBMI(height float64, weight float64) float64 {
	BMI := weight / ((height * height) / 10000)
	return BMI
}
