package main

import (
	"fmt"
	"syscall"
)

func main() {
	readMessage()

}

func readMessage() {
	f, err := syscall.Open("sadsec.txt", syscall.O_RDONLY, 333)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 20)
	nbytes, err := syscall.Read(f, buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf[:nbytes]))
}
