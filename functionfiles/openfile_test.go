package asciiart

import (
	"path/filepath"
	"testing"
)

func TestOpenFile(t *testing.T) {
	// test a valid file path :shadow.txt, standard.txt, thinkertoy.txt
	filePath := "testfiles/shadow.txt"
	_, err := OpenFile(filePath)
	if err != nil {
		t.Errorf("Failed to open and read existing file: %v", err)
	}
}

func TestNonExistentFile(t *testing.T) {
	// test a non existent file
	filePath := "non-existent.txt"
	_, err := OpenFile(filePath)
	if err == nil {
		t.Errorf("Expected error for non existent file, but got none")
	}
}

func TestWrongExtension(t *testing.T) {
	// file will only handle files with .txt extension
	filePath := "testfiles/image.png"
	if filepath.Ext(filePath) == ".txt" {
		t.Errorf("error: file has .txt expected wrong file path")
	}
}

func TestEmptyFile(t *testing.T) {
	filePath := "testfiles/empty.txt"
	_, err := OpenFile(filePath)
	if err != nil {
		t.Errorf("Failed to handle an empty file")
	}
}
