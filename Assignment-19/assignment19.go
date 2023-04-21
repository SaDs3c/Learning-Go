package main

import (
	"fmt"
	"os"
)

type About struct {
	name string
	id   int
}

func main() {
	var ptr *[10]About

	ptr = new([10]About)
	if ptr == nil {
		fmt.Println("Memory could not be allocated.")
		os.Exit(1)
	}

	fmt.Println("Memory was successfully allocated")

}
