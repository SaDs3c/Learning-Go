package main

import "fmt"

type example struct {
	integer int
}

func main() {
	var ptr *example
	var test example

	ptr = &test

	test.integer = 5
	fmt.Printf("%d\n", test.integer)

	(*ptr).integer = 6
	fmt.Printf("%d\n", test.integer)

	ptr.integer = 7
	fmt.Printf("%d\n", test.integer)

}
