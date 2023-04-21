package main

import (
	"fmt"
)

//declare constant vars for these values
const (
	games   int = 4
	players int = 5
)

func main() {
	//declare two dimensional array to store scoreos for 5 players over 4 games
	//declare float array to store our totals for each player later
	var score int
	var scores [4][5]int
	var totals [5]int
	var highestTotal int

	//we need an outer for loop counter and an inner for loop counter
	var gameCounter, playerCounter int

	//start outer loop (4 games)
	for gameCounter = 0; gameCounter < games; gameCounter++ {
		fmt.Printf("Game #%d\n", gameCounter+1)

		//starting inner loop for each player (5 playerss)
		for playerCounter = 0; playerCounter < 5; playerCounter++ {
			fmt.Printf("Enter scoring total for player #%d: ", playerCounter+1)
			//get the score
			fmt.Scanln(&score)

			//put the score into the scores array
			scores[gameCounter][playerCounter] = score
		}
	}

	//create a loop to iterate through each player game by game to create their point totals
	for playerCounter = 0; playerCounter < players; playerCounter++ {
		playerTotal := 0
		for gameCounter = 0; gameCounter < games; gameCounter++ {
			//this player's total will be added to by each game total for the player's in the scores array
			playerTotal += scores[gameCounter][playerCounter]
		}

		//we begin to populating the total array we initialized earlier with each player's total before moving onto the next player
		totals[playerCounter] = playerTotal
	}

	//now we'll create the highest total of 0 to start, and if we iterate through the totals array
	//each total in the array that's higher than the previous highest, will become the new highest
	var playerID int
	for playerCounter = 0; playerCounter < players; playerCounter++ {
		if totals[playerCounter] > highestTotal {
			highestTotal = totals[playerCounter]

			///this playerID variable wil help us keep track of who the current leader is in points
			playerID = playerCounter + 1
		}
	}

	//initialize a new float variable which will be our average
	//use of the (float) operator to change the integer variable being divided into a float operation
	var ppg = float64(highestTotal) / 4
	fmt.Printf("Player #%d had the highest scoring average at %.2f points per game. \n", playerID, ppg)

}
