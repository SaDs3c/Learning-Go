# Assignment 26
We are going to create a Go program that binds a command shell to a listening port, also known as a bind shell. (in TCP)


A good resource through this exercise will be to google "go network programming"

@ftsog also recommend the following great resource:
https://pkg.go.dev/net

# Getting started

# packages
We're going to need to import the following packages
import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
)


Pull these packages up and read through them, learn what they're for and why we need them. you should check out the official doc for all of these packages.

# Summary
If this is your first time working on network programming with Go, you will probably spend some time researching, that is perfectly fine and intendend. Don't get frustrated, if you run into any problem reach out. 
