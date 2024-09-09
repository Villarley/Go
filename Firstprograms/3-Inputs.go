package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	message := greeting(name) // Changed to a variable assignment
	fmt.Print(message)
}

func greeting(name string) string { // Specified return type 'string'
	return fmt.Sprintf("Hello, %s!\n", name) // Corrected string formatting
}
