package asciiart

import (
	"fmt"
	"strings"
	"unicode"

	asciiart "asciiart/mapfile"
)

func DisplayArt(filePath, input string) (string, error) {
	var buildStr strings.Builder // saves on memory
	// fetch the map and handle any errors
	Asciimap, err := asciiart.MapFile(filePath)

	if len(Asciimap) == 0 {
		fmt.Println("Error: Character map is empty, please provide valid char map")
	}
	if err != nil {
		return buildStr.String(), err
	}
	// counter for new lines or no of words split by new
	count := 0
	words := strings.ReplaceAll(input, "\\n", "\n")
	wordsSlice := strings.Split(words, "\n")

	for _, word := range wordsSlice {
		if word == "" {
			count++
			if count < len(wordsSlice) {
				buildStr.WriteString("\n")
			}
		} else {
			// printing the art for each word/substring . Check if each char exists in the map
			for i := 0; i < 8; i++ {
				for _, c := range word {
					s, ok := Asciimap[c]
					if ok {
						buildStr.WriteString(s[i])
					} else {
						buildStr.WriteString("?")
					}

				}
				buildStr.WriteString("\n") // print new line for the next line of art
			}
		}
	}
	return buildStr.String(), nil
}

func IsPrintable(input string) bool {
	for i, val := range input {
		if i < len(input)-1 && input[i] == '\\' {
			if (input[i+1] == 't' || input[i+1] == 'a' || input[i+1] == 'b' || input[i+1] == 'v' || input[i+1] == 'f' || input[i+1] == 'r') && !(input[i+1] == 'n') {
				return false
			}
		}
		if !unicode.IsPrint(val) {
			return false
		}
	}
	return true
}
