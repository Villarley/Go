package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Person struct {
	id     string
	Name   string
	Age    int
	City   string
	Sex    string
	Google string
}

var People []Person

func main() {
	for {
		var selectedOption = askUserOption()
		if selectedOption == 5 {
			break
		}
		performAction(selectedOption)
		fmt.Println(People)
	}
}
func askUserOption() int {
	var option int
	var input string
	for {
		fmt.Println("Enter the option you want to perform:\n1. Add User\n2.Edit User\n3.See all the users\n4.Delete User\n5.Exit")
		fmt.Scanln(&input)

		var err error
		option, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
		} else {
			break // Break the loop
		}
	}
	return option
}
func performAction(selectedOption int) {
	var userId string
	switch selectedOption {
	case 1:
		addUser()
	case 2:
		userId = obtainUserId()
		editUser(userId)
	case 3:
		showAllUsers()
	case 4:
		userId = obtainUserId()
		deleteUser(userId)
	case 5:
		break
	default:
		fmt.Println("Not valid")
	}
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
func addUser() {
	var newPerson Person
	newPerson.id = generateRandomCode(6)
	fmt.Println("Enter the name of the new user:")
	fmt.Scanln(&newPerson.Name)
	fmt.Println("Enter the age of the new user:")
	fmt.Scanln(&newPerson.Age)
	fmt.Println("Enter the city of the new user:")
	fmt.Scanln(&newPerson.City)
	fmt.Println("Enter the sex of the new user:")
	fmt.Scanln(&newPerson.Sex)
	fmt.Println("Enter the google of the new user:")
	fmt.Scanln(&newPerson.Google)
	People = append(People, newPerson)
}
func editUser(id string) {
	for i := range People {
		if People[i].id == id {
			fmt.Println("Enter the name of the user:")
			fmt.Scanln(&People[i].Name)
			fmt.Println("Enter the age of the user:")
			fmt.Scanln(&People[i].Age)
			fmt.Println("Enter the city of the user:")
			fmt.Scanln(&People[i].City)
			fmt.Println("Enter the sex of the user:")
			fmt.Scanln(&People[i].Sex)
			fmt.Println("Enter the google of the user:")
			fmt.Scanln(&People[i].Google)
			break
		} else {
			continue
		}
	}
}
func showAllUsers() {
	for i := range People {
		fmt.Println(People[i])
	}
}
func deleteUser(id string) {
	var pos int
	for i := range People {
		if People[i].id == id {
			pos = i
			if pos < 0 || pos >= len(People) {
				return
			}

			People = append(People[:pos], People[pos+1:]...)
			break
		} else {
			continue
		}
	}
}
func obtainUserId() string {
	var id string
	for {
		fmt.Println("Enter the id of the user you want to edit:")
		fmt.Scanln(&id)
		if len([]rune(id)) == 6 {
			break
		} else {
			fmt.Println("Please enter a valid id")
			obtainUserId()
		}
	}
	return id
}
