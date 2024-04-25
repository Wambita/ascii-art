package asciiart

import (
	"fmt"
	"testing"
)

func TestMapFile(t *testing.T) {
	m, err := MapFile("../standard.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println("A:", m['A'])
	for _, line := range m['N'] {
		fmt.Println(line)
	}

}

