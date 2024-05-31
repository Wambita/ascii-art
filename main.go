package main

import (
	"fmt"
	"os"

	asciiart "asciiart/functionfiles"
)

func main() {
	var filePath string
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . \"something\" <BANNER>")
	}

	if len(os.Args) == 2 {
		filePath = "standard.txt"
	}

	if len(os.Args) == 3 {
		flag := os.Args[2]
		filePath = "standard.txt"
		switch flag {
		case "shadow", "sh", "Sh", "SHADOW":
			filePath = "shadow.txt"
		case "thinkertoy", "t", "T", "THINKERTOY":
			filePath = "thinkertoy.txt"
		case "standard", "Standard", "S", "s", "STANDARD":
			filePath = "standard.txt"
		default:
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . \"something\" <BANNER>")
			return

		}

	}
	input := os.Args[1]
	s, err := asciiart.DisplayArt(filePath, input)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Print(s)
}
