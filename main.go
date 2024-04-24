package main

// get the file
// create a map function to map the ascii chars to the art in the banner
// func to convert the input string to ascii
// func to print the input in the form of ascii art
// handle possible scenarios, new lines , spaces , chars that are not in the ascii table
// handle errors
// test the program

// get the file
// os / Readfile/ os.open then scanner.

import (
	"fmt"
	"log"

	asciiart "asciiart/openFile"
)

func main() {
	filePath := "shadow.txt"
	content, err := asciiart.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}
