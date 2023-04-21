package main

import (
	"fmt"
)

//define a value for PIE
const (
	PIE float64 = 3.14
)

func main() {
	//initilize variable
	var radius float64

	//get user input and store it
	fmt.Print("Enter the radius of your circle: ")
	_, err := fmt.Scanf("%f", &radius)
	if err != nil {
		fmt.Println(err)
		return
	}

	//do the maths
	var area float64
	area = PIE * (radius * radius)

	fmt.Println("The Area of your Circle is", area)
}
