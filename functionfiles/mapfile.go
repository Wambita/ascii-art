package asciiart

import "fmt"

func MapFile(filePath string) (map[rune][]string, error) {
	asciiMap := make(map[rune][]string)
	currentChar := ' '
	// filePath := "standard.txt"
	lines, err := OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		if line == "" {
			continue
		}
		if len(asciiMap[currentChar]) == 8 {
			currentChar++
		}
		asciiMap[currentChar] = append(asciiMap[currentChar], line)

	}

	// validation for the ascii map
	if len(lines) != 855 || len(asciiMap) == 0 || len(asciiMap) != 95 {
		return nil, fmt.Errorf("the banner file %q does not contain expected format", filePath)
	}
	return asciiMap, nil
}
