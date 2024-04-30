package main

import (
	"fmt"
	"os"

	asciiart "asciiart/displayart"
)

func main() {
	var filePath string
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println(`usage: go run . "input string" or usage: go run . "input string " -font`)
		return
	}

	if len(os.Args) == 2 {
		filePath = "standard.txt"
	}

	if len(os.Args) == 3 {
		flag := os.Args[2]
		switch flag {
		case "-shadow", "-s", "-sw":
			filePath = "shadow.txt"
		case "-thinkertoy", "-t", "T":
			filePath = "thinkertoy.txt"
		default:
			fmt.Println(`usage: go run . "input string " -font`)
			return

		}

	}
	input := os.Args[1]
	// check if char is printable
	if !asciiart.IsPrintable(input) {
		fmt.Println("unprintable characters present")
	}
	s, err := asciiart.DisplayArt(filePath, input)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Print(s)
}
