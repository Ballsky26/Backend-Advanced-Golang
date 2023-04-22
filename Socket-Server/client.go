package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":")
	if err != nil {
		fmt.Println("Error dial...")
	}
	defer connection.Close()

	_, err = connection.Write([]byte("Hello Server! Ballsky"))

	// Read message from server
	buffer := make([]byte, 1024)
	mlen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error read...")
	}
	fmt.Println("Received : ", string(buffer[:mlen]))
}
