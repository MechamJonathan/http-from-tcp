package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Could not open %s: %s\n", inputFilePath, err)
	}
	defer file.Close()

	fmt.Printf("Reading data from: %s\n", inputFilePath)
	fmt.Println("------------------------------------")

	var currentLine string
	for {
		buf := make([]byte, 8)
		n, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if currentLine != "" {
					fmt.Printf("read: %s\n", currentLine)
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
				fmt.Printf("read: %s\n", completeLine)
				currentLine = ""
			}
		}
	}
}
