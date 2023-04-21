package main

import (
	"fmt"
)

func main() {
	var numerator int
	var denominator int
	var remainder int

	fmt.Print("Enter a numerator: ")
	_, err := fmt.Scanln(&numerator)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter a denomenator: ")
	_, err = fmt.Scanln(&denominator)
	if err != nil {
		fmt.Println(err)
		return
	}

	remainder = numerator % denominator
	if remainder != 0 {
		fmt.Println("There is a remainder!")
		return
	}

	fmt.Println("There is NOT a remainder!")

}
