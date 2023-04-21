# Assignment 23
Create a simple structure of your choice. Mine will have a single member, an int called integer. Do not create an instance of the struct yet. Declare a pointer variable that points to your struct. This pointer doesn't actually point to anything yet because there are no instance of the first struct. Create an instance of the struct. Now initialize your pointer to point to the struct instance.

Finally, assign values to the integer member of your original struct in three ways:
* assign a value by using dot-notation for your instance and then print the values
* assign a value by using dot-notation for the pointer initialize to the instance and then print the value
* use shorthand of the second way

# Psudo Code
type example struct {
	integer int
}

var ptr *example
var test example

<pointer initializatin statement>

test.integer = x
print

(*ptr).integer = y is the same thing as ptr.integer = y (shorthand, pointer is automatically dereference)
print




# Example Output
ftsog@nigeria:~/LearningGo/Assignment-23$ ./assignment23
5
6
7