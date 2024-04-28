package asciiart

import (
	"fmt"
	"testing"
)

//checking if the MapFile function maps the runes correctly to their corresponding art 
func TestMapFile(t *testing.T) {
	m, err := MapFile("../test.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range m['A'] {
		fmt.Println(line)
	}
}
