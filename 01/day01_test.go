package main

import (
	"testing"
)

func TestMainOneStar(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
			expected: 24000,
		},
	}
	for _, test := range tests {
		byteSlice := []byte(test.input)
		result := MainOneStar(byteSlice)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
