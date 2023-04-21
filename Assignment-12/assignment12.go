package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var (
		number  int
		counter int
		head    int
		tail    int
		seed    = time.Now().UnixNano()
	)

	fmt.Print("How many times would you like to flip the coin? ")
	fmt.Scanln(&number)

	counter = number
	rand.Seed(seed)
	for {
		if counter <= 0 {
			break
		}

		res := rand.Intn(2)
		if res == 0 {
			head += 1
		} else if res == 1 {
			tail += 1
		}
		counter -= 1
	}

	fmt.Println(head, "heads")
	fmt.Println(tail, "tails")

}
