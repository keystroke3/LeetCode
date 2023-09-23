package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			input:    []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			expected: []int{0, 0, 0, -36, 0, 0, 0},
		},
		{
			input:    []int{10, 15, 20, 25},
			expected: []int{7500, 5000, 3750, 3000},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			prods := productExceptSelf(test.input)
			if !reflect.DeepEqual(prods, test.expected) {
				t.Errorf("Expected %v for input %v. Got %v\n", test.expected, test.input, prods)
			}
		})
	}
}
