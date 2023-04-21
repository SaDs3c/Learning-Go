package main

import (
	"fmt"
)

type stats struct {
	points int
	games  int
}

func main() {
	var players [5]stats

	//loop through getting input for both stats members for each stats players struct
	for counter := 0; counter < 5; counter++ {
		fmt.Printf("Enter player %d's point total: ", counter+1)
		fmt.Scanln(&players[counter].points)

		fmt.Printf("Enter player %d's game total: ", counter+1)
		fmt.Scanln(&players[counter].games)
	}

	//you know this by now, you're a Go god
	for counter := 0; counter < 5; counter++ {
		var average float64 = float64(float64(players[counter].points) / float64(players[counter].games))
		fmt.Printf("Player %d's scoring average was %0.2f ppg.\n", counter+1, average)
	}
}
