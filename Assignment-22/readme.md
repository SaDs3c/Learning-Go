# Assignment 22
more struct? God no, please no. I'm sorry, but yes. One more excercise. we are putting array in our structs, we are leveling up.

The goal here is to create a program that takes input for five player's point totals and games played totals and then calculate the scoring averages for each player. Create a struct with members for total points and games played. Create an arrray of structures that is five elements deep (one for each player) and loop through the array filling it with the point totals and games played totals for each player.

This array of structs should look like this in pseudo code:
type stats struct {
	points int
	games int
}

var player [5]stats

you can declare an in counter variable and use it to iterate through the array for both points and games


# Example Output
ftsog@nigeria:~/LearningGo/Assignment-22$ ./assignment22
Enter player 1's point total: 44
Enter player 1's game total: 5
Enter player 2's point total: 21
Enter player 2's game total: 5
Enter player 3's point total: 33
Enter player 3's game total: 5
Enter player 4's point total: 67
Enter player 4's game total: 5
Enter player 5's point total: 8
Enter player 5's game total: 5
Enter player 1's scoring average was 8.80 ppg.
Enter player 2's scoring average was 4.20 ppg.
Enter player 3's scoring average was 6.60 ppg.
Enter player 4's scoring average was 13.40 ppg.
Enter player 5's scoring average was 1.60 ppg.