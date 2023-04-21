package main

import "fmt"

type example struct {
	integer   int
	character rune
	floatie   float64
}

func main() {
	//create a variable of type example
	var sample example
	//use the . notation to refer to the struct member of the 'sample'
	sample.integer = 50
	sample.character = 'W'
	sample.floatie = 3.14

	fmt.Printf("This is the int: %d", sample.integer)
	fmt.Printf(", This is the char: %c", sample.character)
	fmt.Printf(", This is the flaot: %0.2f", sample.floatie)
}
