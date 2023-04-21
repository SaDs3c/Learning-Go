package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) < 3 || len(arg) > 3 {
		fmt.Println("Usage: ./assignment9 Firstname Lastname")
		return
	}

	fmt.Println("Hello,", arg[1], arg[2])

}
