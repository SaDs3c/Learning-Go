# Assignment 18
Compile and run the following code and examine the output

package main

func hello() {
	arr := [13]int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	fmt.Println(arr)
}

func main() {
	hello()
}

Now, compile this code and examine the output:

package main

func hello() {
	arr := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	fmt.Println(arr)
}

func main() {
	hello()
}


what is the difference? 

Extra Credit: Google and learn more about slice