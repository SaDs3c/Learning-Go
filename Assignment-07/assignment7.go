package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Enter the value of variable 'A': ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter the value of variable 'B': ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter the value of variable 'C': ")
	_, err = fmt.Scanln(&c)
	if err != nil {
		fmt.Println(err)
		return
	}

	//start doing the math
	var solution1, solution2 float64
	solution1 = (-b + math.Sqrt((b*b)-(4*a*c))) / (2 * a)
	solution2 = (-b - math.Sqrt((b*b)-(4*a*c))) / (2 * a)

	//plug the solution back to into the equation and check that it's correct: 0 = ax^2 + bx + c
	if ((a * (solution1 * solution1)) + (b * solution1) + c) == 0 {
		fmt.Printf("The solution using the '+' operator is: %0.2f\n", solution1)
	} else {
		fmt.Printf("The solution using the '+' operator is: %0.2f but you might want to double-check that...", solution1)
	}

	if ((a * (solution2 * solution2)) + (b * solution2) + c) == 0 {
		fmt.Printf("The solution using the '-' operator is: %0.2f", solution2)
	} else {
		fmt.Printf("The solution using the '-' operator is: %0.2f but you might want to double-check that...", solution2)
	}
}
