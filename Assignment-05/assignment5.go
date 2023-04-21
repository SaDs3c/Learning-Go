package main

import (
	"fmt"
)

func main() {
	//initilize var
	var seconds float64

	//get value for var
	fmt.Print("Enter the amount of seconds: ")
	_, err := fmt.Scanln(&seconds)
	if err != nil {
		fmt.Println(err)
		return
	}

	//make sure the hours value is a whole number and doesn't consider the
	//decimal by using the (int) operation.
	//for instance 1.2 hours will stored as 1.0 here
	hours := int(seconds / 3600)

	//repeat for mins
	mins := int((int(seconds) - (hours * 3600)) / 60)

	//remainingSeconds is simply what's leftover
	remainingSeconds := int(int(seconds) - (hours * 3600) - (mins * 60))
	fmt.Printf("%v seconds is equal to %v hours, %v minutes, and %v seconds.", seconds, hours, mins, remainingSeconds)
}
