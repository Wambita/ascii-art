package asciiart

import (
	asciiart "asciiart/openfile"
)

func MapFile(filePath string) (map[rune][]string, error) {
	asciiMap := make(map[rune][]string)
	currentChar := ' '
	// filePath := "standard.txt"
	lines, err := asciiart.OpenFile(filePath)
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
	return asciiMap, nil
}
