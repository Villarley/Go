package main

import "fmt"

func main() {
	var askedLength int
	fmt.Println("Enter the quantity of numbers the list will have")
	fmt.Scanln(&askedLength)
	numbers := fillArray(askedLength)
	bubbleSort(numbers)
	fmt.Println(numbers)
}

func fillArray(askedLength int) []int {
	numbers := make([]int, askedLength)
	var number int
	for i := 0; i < askedLength; i++ {
		fmt.Println("Enter a value ")
		fmt.Scanln(&number)
		numbers[i] = number
	}
	return numbers
}
func bubbleSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		// Realizar 'n-i-1' comparaciones en la i-ésima pasada
		for j := 0; j < n-i-1; j++ {
			// Comparar y realizar intercambio si necesario
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
