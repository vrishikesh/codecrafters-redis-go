package main

import (
	"fmt"
	"io"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		fmt.Println("Starting server")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		
	}
}

func handleConnection(c net.Conn) {
	_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing to connection: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Closing connection")
		err = conn.Close()
		if err != nil {
			fmt.Println("Error closing connection: ", err.Error())
			os.Exit(1)
		}

	defer c.Close()
	io.Copy(c, c)
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
