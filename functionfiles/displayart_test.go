package asciiart

import (
	"os"
	"testing"
)

func readFile(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(data)
	//
}

func TestDisplayArt(t *testing.T) {
	test1 := []struct {
		input    string
		expected string
	}{
		{
			"hello There 1 to 2!",
			readFile("testfiles/test1.txt"),
		},
	}
	for _, test := range test1 {
		filePath := "../standard.txt"
		output, err := DisplayArt(filePath, test.input)
		if err != nil {
			t.Errorf("Error %s", err)
		}
		os.WriteFile("testfiles/test2.txt", []byte(output), 0o644)
		if output != test.expected {
			os.WriteFile("test3.txt", []byte(output), 0o644)
			t.Errorf("DisplayArt(%s), expexted(\n%s\n), got(\n%s\n)", test.input, test.expected, output)
		}
	}
}
