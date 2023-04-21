# Assignment 15
* Create a int variable with var variable int
* Create variable that is a pointer to the int variable with var prt *int = &variable

Use the following variable references to replace the x in the appropriate print statement below
* *pointer
* variable
* pointer
* &variable

Print statements:
* fmt.Printf("The value of int variable is: %d\n", x)
* fmt.Printf("The value of the pointer to the int variable is %d\n", x)
* fmt.Printf("The memory address of the int variable is: %d\n", x)
* fmt.Printf("The value held at the memory location that the pointer is pointing to is: %%d\n", x)

Hint: Google and read about pointers

Extra Credit: There's more than one set of correct answer! Find a different way to get the same output

# Example Output
ftsog@nigeria:~/LearningGo/Assignment-15$ ./assignment15
The value of int variable is: 15
The value of the pointer to the int variable is: 0xc0000141c8
The memory address of the int variable is: 0xc0000141c8
The value held at the memory location that the pointer is pointing to is: 15