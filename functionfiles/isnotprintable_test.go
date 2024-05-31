package asciiart

import "testing"

func TestIsNotPrintable(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected bool
	}{
		{name: "Printable Input", str: "Hello World", expected: false},
		{name: "Printable Input", str: "Hello\nWorld", expected: false},
		{name: "Printable Input", str: "Hello\\t World", expected: true},
		{name: "Printable Input", str: "Hello\t World", expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := IsNotPrintable(test.str)
			if output != test.expected {
				t.Errorf("Error: IsNotPrintable(%v) \nExpected: %v \nGot: %v", test.str, test.expected, output)
			}
		})
	}
}
