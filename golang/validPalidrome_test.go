package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "race car",
			expected: true,
		},
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "  ",
			expected: false,
		},
		{
			input:    " ",
			expected: false,
		},
		{
			input:    "aa",
			expected: true,
		},
		{
			input:    " man, a plan, a canal: Panama",
			expected: false,
		},

		{
			input:    "race a car",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			palindrome := isPalindrome(test.input)
			if palindrome != test.expected {
				t.Errorf("Expect %v for input %v but got %v\n", test.expected, test.input, palindrome)
			}
		})
	}
}
