package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Enter a number between 1 and 500: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)
		return
	}

	if number < 1 || number > 500 {
		fmt.Println("Your number was not in any of our ranges.")
		return
	}

	if number >= 1 && number <= 100 {
		fmt.Println("Your number is between 1 and 100!")
		return
	}

	if number >= 101 && number <= 200 {
		fmt.Println("Your number is between 101 and 200!")
		return
	}

	if number >= 201 && number <= 300 {
		fmt.Println("Your number is between 201 and 300!")
		return
	}

	if number >= 301 && number <= 400 {
		fmt.Println("Your number is between 301 and 400!")
		return
	}

	if number >= 401 && number <= 500 {
		fmt.Println("Your number is between 401 and 500!")
		return
	}

}
