package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("testfile.txt", os.O_WRONLY|os.O_CREATE, 700)
	if err != nil {
		fmt.Println("The open operation failed...")
		os.Exit(1)
	} else {
		fmt.Println("Then open operation succeeded!")
	}

	_, err = file.Write([]byte("Writing test data to the file."))
	if err != nil {
		fmt.Println("The write operation failed...")
		os.Exit(1)
	} else {
		fmt.Println("The write operation succeeded!")
	}
}
