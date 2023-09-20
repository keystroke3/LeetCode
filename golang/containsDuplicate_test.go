package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			input:    []int{1, 1, 1, 2, 3, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{-1, 2, 3, 1},
			expected: false,
		},
		{
			input:    []int{-1, -50, -50, 2, 3, 1},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {

			_, duped := containsDuplicate(test.input)
			if duped != test.expected {
				t.Errorf("Test failed, expected %v for input %v, but got %v\n", test.expected, test.input, duped)
			}
		})
	}
}
