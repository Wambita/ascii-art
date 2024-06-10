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
			if fileName == "standard.txt" || fileName == "shadow.txt" || fileName == "thinkertoy.txt" {
				fmt.Println("File does not exist:", fileName)
				GetFile(fileName)
				return nil, err
			}
			fmt.Println("invalid file name used")
			return nil, err

		} else {
			return nil, err
		}
	}

	if filepath.Ext(fileName) != ".txt" {
		fmt.Println("Error: wrong file extension. use extension .txt")
		return nil, err
	}
	fileinfo, err := os.Stat(fileName)
	if fileinfo.Size() == 0 {
		return nil, fmt.Errorf("%q is empty", fileName)
	}
	if err != nil {
		return nil, fmt.Errorf("error: %q", err)
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
