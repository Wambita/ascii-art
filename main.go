package main

import (
	"fmt"
	"os"

	asciiart "asciiart/displayart"
)

func main() {
	filePath := "standard.txt"
	if len(os.Args) != 2 {
		fmt.Println(`usage: go run . "input string"`)
		os.Exit(1)
	}
	input := os.Args[1]
	s, err := asciiart.DisplayArt(filePath, input)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Print(s)
}
