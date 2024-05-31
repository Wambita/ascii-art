package asciiart

import (
	"fmt"
	"strings"
)

func IsNotPrintable(str string) bool {
	NonPrintableChars := []string{"\\a", "\\b", "\\t", "\\v", "\\f", "\\r", "\a", "\b", "\t", "\v", "\f", "\r"}

	for _, char := range NonPrintableChars {
		if contains := strings.Contains(str, char); contains {
			fmt.Println("Error: input contains non-printable chars:", char)
			return true
		}
	}
	return false
}
