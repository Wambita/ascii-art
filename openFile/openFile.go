package asciiart

// Function that takes in filepath as an argument and returns contents of file as a string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func OpenFile(filePath string) (string, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("error: file does not exist")
		}
		return "", fmt.Errorf("error: error checking file")
	}

	// open the file
	ArtFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer ArtFile.Close()

	// create scanner object to read the file and return as string
	scanner := bufio.NewScanner(ArtFile)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	content := strings.Join(lines, "\n")
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}
