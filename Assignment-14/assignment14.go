package main

import "fmt"

func main() {
	var number int = 21
	var ptr *int = &number

	fmt.Printf("The value of out pointer is %p\n", ptr)
}
