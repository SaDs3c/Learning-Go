package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
)

const (
	network   string = "tcp"
	port      string = "9000"
	host      string = "localhost"
	seperator string = ":"
)

func main() {
	fmt.Println("Listening on tcp port", port)
	listener, err := net.Listen(network, host+seperator+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		Handler(conn)
	}

}

func Handler(conn net.Conn) {
	command, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("Error: Invalid Input %v\n", err)))
	}

	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("Error: Invalid command %v\n", err)))
	}

	conn.Write(output)

	conn.Close()
}
