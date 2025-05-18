package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(c net.Conn) <-chan string {
	linesChan := make(chan string)

	go func() {
		defer c.Close()

		var currentLine string
		for {
			buf := make([]byte, 8)
			n, err := c.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					if currentLine != "" {
						linesChan <- currentLine
					}
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}

			parts := strings.Split(string(buf[:n]), "\n")

			for i, part := range parts {
				if i == len(parts)-1 {
					currentLine += part
				} else {
					completeLine := currentLine + part
					linesChan <- completeLine
					currentLine = ""
				}
			}
		}

		close(linesChan)
	}()

	return linesChan
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Printf("error: %s", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error: %s", err)
			continue
		}

		fmt.Println("connection has been accepted")

		linesChan := getLinesChannel(conn)
		for line := range linesChan {
			fmt.Println(line)
		}
		fmt.Println("connection has been closed")
	}
}
