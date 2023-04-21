package main

import (
	"fmt"
)

func main() {
	var score int
	var play string
	var totalScore int
	var counter int

	for {
		fmt.Print("Enter a test score: ")
		_, err := fmt.Scanln(&score)
		if err != nil {
			fmt.Println(err)
			return
		}

		counter += 1

		totalScore += score
		fmt.Print("Would you like to continue? Y/N ")
		_, err = fmt.Scanln(&play)
		if err != nil {
			fmt.Println(err)
			return
		}

		if play == "Y" {
			continue
		} else if play == "N" {
			break
		} else {
			fmt.Println("Wrong input")
			return
		}

	}

	avgScore := float64(totalScore) / float64(counter)
	fmt.Printf("%0.2F is the average", avgScore)
}
