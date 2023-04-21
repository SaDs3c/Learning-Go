package main

import "fmt"

func main() {
	var array [10]int

	fmt.Printf("%p\n", &array)
	fmt.Printf("%p", &array[0])

}
