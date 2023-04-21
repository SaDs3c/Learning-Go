# Assignment 24
For this exercise we're going to build a Go program that executes syscalls! We've already learned basically everything we need to know in Go up until this point to accomplish this. We are going to keep it simple and we're going to use the Open() syscall only here. it should be noted that technically we are not making syscall directly in user program, but instead we are using method defined in the "os" package for our purpose

# Open() Syscall
Next, we need to determine how to open the file using the os.OpenFile(). To do this, we can consult official os package doc. The method definition according to the officail doc is: func OpenFile(name string, flag int, perm Filemode) (*File, error)
* name string is the relative path or absolute path to the file that is to be open
* flag int is "A bitwise 'or' seperated list of values that determine the method in which the file is to be opened(whether it should be read only, read/write, whether it should be cleared when opened, etc)"

These are legal list of values for flag int

* O_RDONLY: read-only opens the file for reading only.
* O_WRONLY: write-only opens the file for writing only.
* O_RDWR: read-write opens the file for both reading and writing.
* O_APPEND: add the data to the end of the file while writing to the file
* O_CREATE: when opening the file, if there is no file with the specified name, it creates
* O_EXCL: Used with O_CREATE. A fiel with the specified name must not exist in order to create a new file
* O_SYNC: Allows the file to be opened for synchronous I/O
* O_TRUNC: initially clear all data from the file

The first thing to realize is that our file doesn't exist. We will definitively be needing the O_CREATE value. This method also specifies we need a third parameter. the third parameter is perm *FileMode.

A FileMode represents a file's mode and permission bit


so, this perm *FileMode parameter is going to be added. we'll set permission on the file to 700, so we'll use the 700 value for the perm FileMode parameter.

we'll open our file in write mode (O_WRONLY), so this is our code.

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("testfile.txt", os.O_WRONLY|os.O_CREATE, 700)
	if err != nil {
		fmt.Println("The open operation failed...")
		os.Exit(1)
	}

	_ = file

	fmt.Println("Then open operation succeeded!")
}

We've replaced name string with the name of our file and we've replaced flag int with O_WRONLY and O_CREATE seperated by the logical or operator |. Lastly, we've  placed 700 as the value for perm FileMode parameter.

we set two variable to hold the values returned by the OpenFile method, the mehod returned values are of type *os.file and error, the former handle the rest of the file I/O operation and the latter contain data concerning error that occured when the OpenFile Method get Called.

let's run it and check the output


# Example Output
ftsog@nigeria:~/LearningGo/Assignment-24$ ./assignment24
Then open operation succeeded!

Awesome, our code apparently worked. let's check the directory and see the permission for the newly created testfile.txt

ftsog@nigeria:~/LearningGo/Assignment-24$ ls -lah
total 4.0K
drwxrwxrwx 1 ftsog ftsog 4.0K Apr 18 15:01 .
drwxrwxrwx 1 ftsog ftsog 4.0K Apr 17 17:46 ..
-rwxrwxrwx 1 ftsog ftsog  265 Apr 18 15:01 assignment24.go
-rwxrwxrwx 1 ftsog ftsog 3.5K Apr 18 14:58 readme.md
-rwxrwxrwx 1 ftsog ftsog    0 Apr 18 15:01 testfile.txt
