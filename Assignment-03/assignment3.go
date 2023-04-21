package main

import (
	"fmt"
)

func main() {

	//initilize vars
	var firstName string
	var lastName string

	//prompt user to input first and last name and use fmt.Scanf() to store those to the initilize vars
	fmt.Print("Enter first name: ")
	_, err := fmt.Scanln(&firstName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter Last Name: ")
	_, err = fmt.Scanln(&lastName)
	if err != nil {
		fmt.Println(err)
		return
	}

	//print the welcome message!
	fmt.Println("Hello", firstName, lastName+"!")
}
