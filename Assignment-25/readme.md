# Assignment 25
This assignment will pick up from where we left off. we have successfully implement an Open() syscall using the OpenFile() method and will now pair that with a Write() syscall to insert text into our open file.

# Getting Started
Again we will refernce the doc site.

# Write() Syscall
jumping right in, the official doc gives us the following Method definition: func (f *File) Write(b []bytes) (n int, err error)
it elaborate on the methods by saying -> Write writes len(b) bytes from b to the File. it returns the number of bytes written and an error, if any. Write returns a non-nil error when n != len(b).

message := []byte("Write test data to the file")
the length of our message is 30 to calculate this let's just use python in the terminal real quick
ftsog@nigeria:~/LearningGo/Assignment-25$ python
Python 3.10.8 (main, Oct 24 2022, 10:07:16) [GCC 12.2.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> string = "Writing test data to the file."
>>> print(len(string))
30
>>>

ok, so our message is 30 bytes long. At this point, we have everything we need to make a Write() call, our code now look like this

package main

func main() {
	
	file, err := os.OpenFile("testfile.txt", os.O_WRONLY|os.O_CREATE, 700)
	if err != nil {
		fmt.Println("The open operation failed...")
		os.Exit(1)
	} else {
		fmt.Println("Then open operation succeeded!")
	}

	_, err = file.Write([]byte("Writing test data to the file."))
	if err != nil {
		fmt.Println("The write operation failed...")
		os.Exit(1)
	} else {
		fmt.Println("The write operation succeeded!")
	}
	
}


ftsog@nigeria:~/LearningGo/Assignment-25$ ./assignment25
Then open operation succeeded!
The write operation succeeded!
ftsog@nigeria:~/LearningGo/Assignment-25$ cat testfile.txt
Writing test data to the file.

Now we just need to add a Close() operation to our file so that our file is not left open, I'will leave this exercise to you!