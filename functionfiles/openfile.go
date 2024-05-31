package asciiart

// Function that takes in filepath as an argument and returns contents of file as a string

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func OpenFile(fileName string) ([]string, error) {
	// open the file
	ArtFile, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist:", fileName)
			ArtFile = GetFile(fileName)
			os.Exit(0)
		} else {
			return nil, err
		}
	}

	if filepath.Ext(fileName) != ".txt" {
		fmt.Println("Error: wrong file extension. use extension .txt")
		return nil, err
	}

	defer ArtFile.Close()

	// create scanner object to read the file and return as string
	scanner := bufio.NewScanner(ArtFile)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	// content := strings.Join(lines, "\n")
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
