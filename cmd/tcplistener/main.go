package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MechamJonathan/http-from-tcp/internal/request"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Printf("error: %s", err)
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error: %s", err)
			continue
		}

		fmt.Println("connection has been accepted")

		req, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalf("error parsing request: %s\n", err.Error())
		}
		fmt.Println("Request line:")
		fmt.Printf("- Method: %s\n", req.RequestLine.Method)
		fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)
		fmt.Println("Headers:")
		for key, value := range req.Headers {
			fmt.Printf("- %s: %s\n", key, value)

		}
	}
}
