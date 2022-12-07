package main

import (
	"fmt"

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
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			break
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 4096)

	for {
		n, er := c.Read(buf)
		if n > 0 {
			// _, ew := c.Write(buf[:n])
			_, ew := c.Write([]byte("+PONG\r\n"))
			if ew != nil {
				break
			}
		}
		if er != nil {
			break
		}
	}

	// defer c.Close()
	// io.Copy(c, c)
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
