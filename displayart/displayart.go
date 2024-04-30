package asciiart

import (
	"strings"

	asciiart "asciiart/mapfile"
)

func DisplayArt(filePath, input string) (string, error) {
	var b strings.Builder // saves on memory
	// fetch the map and handle any errors
	m, err := asciiart.MapFile(filePath)
	if err != nil {
		return b.String(), err
	}
	// counter for new lines or no of words split by new
	count := 0
	words := strings.Split(input, "\\n")
	for _, word := range words {
		if word == "" {
			count++
			if count < len(words) {
				b.WriteString("\n")
			}
			continue
		}
		// printing the art for each word/substring . Check if each char exists in the map
		for i := 0; i < 8; i++ {
			for _, c := range word {
				s, ok := m[c]
				if ok {
					b.WriteString(s[i])
				} else {
					b.WriteString("?")
				}

			}
			b.WriteString("\n") // print new line for the next line of art
		}

	}
	return b.String(), nil
}

func IsPrintable(input string) bool {
	for i := 0; i < len(input); i++ {
		c := input[i]
		if !(c < ' ' || c > '~' || c == '\n') {
			// Return false if the character is out of the basic printable range
			return false
		}
		if c == '\\' {
			if i+1 < len(input) && input[i+1] == '\\' {
				// Skip the next character if it is a valid escaped backslash
				i++ // Skip the escaped backslash
			} else {
				// If it's not a valid escape (like \n, \t, etc.), and also not another backslash, it's invalid
				return false
			}
		}
	}
	return true
}
