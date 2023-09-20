package main

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		input    [2]string
		expected bool
	}{
		{
			input:    [2]string{"anagram", "nagaram"},
			expected: true,
		},
		{
			input:    [2]string{"anagram", "nnnnagaram"},
			expected: false,
		},
		{
			input:    [2]string{"dream", "armored"},
			expected: false,
		},
		{
			input:    [2]string{"night", "thing"},
			expected: true,
		},
		{
			input:    [2]string{"nights", "thing"},
			expected: false,
		},
		{
			input:    [2]string{"the meaning of life", "the fine game of nil"},
			expected: true,
		},
		{
			input:    [2]string{"bottoms up!", "pubs motto!"},
			expected: true,
		},
		{
			input:    [2]string{"mississippi", "ipsismisisp"},
			expected: true,
		},
		{
			input:    [2]string{"abcde", "adc be"},
			expected: true,
		},
		{
			input:    [2]string{"ab", "a"},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			anagram := isAnagram(test.input[0], test.input[1])
			if anagram != test.expected {
				fmt.Printf("Test failed, expected %v for input '%v, %v' but got %v\n", test.expected, test.input[0], test.input[1], anagram)
			}
		})
	}
}
