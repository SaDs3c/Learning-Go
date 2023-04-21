package main

import "fmt"

func main() {
	//declare int variable and a pointer to the int varaiable
	var variable int = 15
	var ptr *int = &variable

	//standard, just printing the value
	fmt.Printf("The value of int variable is: %d\n", variable)

	//since we assign the pointer to 'variable', out pointer's value is now the memory address of 'variable'
	fmt.Printf("The value of the pointer to the int variable is: %p\n", ptr)

	//using the ampersand before variable means we're printing it's memory address instead of it's value
	fmt.Printf("The memory address of the int variable is: %p\n", &variable)

	//using the asterisk means we will print the value held at the memory address pointed to by the pointer
	//this process is called dereferencing
	fmt.Printf("The value held at the memory location is pointing to is: %d\n", *ptr)
}
