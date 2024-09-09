package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var names = fillArrayWithNames()
	sort.Strings(names)
	fmt.Println(names)
}
func fillArrayWithNames() []string {
	var askedLength int
	var input string

	for {
		fmt.Println("Enter the amount of friends (a number):")
		fmt.Scanln(&input)

		// Intentar convertir la entrada en un número
		var err error
		askedLength, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
		} else {
			break // Salir del bucle si la entrada es válida
		}
	}

	names := make([]string, 0, askedLength)
	var name string

	for i := 0; i < askedLength; i++ {
		fmt.Println("Enter the name of a friend:")
		fmt.Scanln(&name)
		names = append(names, name)
	}
	return names
}
