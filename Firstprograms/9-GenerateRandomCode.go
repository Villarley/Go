package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var askedLength int
	var randomCode string
	askedLength = askLength()
	randomCode = generateRandomCode(askedLength)
	fmt.Println(randomCode)
}

func askLength() int {
	var input string
	var askedLength int
	for {
		fmt.Println("Enter the length of the code")
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
	return askedLength
}
func generateRandomCode(length int) string {
	var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}
