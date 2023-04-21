package main

import (
	"fmt"
)

func main() {

	//initilize vars
	var first string
	var last string

	//prompt user to input first and last name and use fmt.Scanf() to store those to the initilize vars
	fmt.Print("Enter your first name: ")
	_, err := fmt.Scanf("%s \n", &first)
	if err != nil {
		fmt.Printf(err)
		return
	}

	fmt.Print("Enter your last name: ")
	_, err = fmt.Scanf("%s", &last)
	if err != nil {
		fmt.Println(err)
		return
	}

	//print the welcome message!
	fmt.Printf("Hello %s %s!\n", first, last)
}
