package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected [2]int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: [2]int{0, 1},
		},
		{
			nums:     []int{10131, 6, 8, 10, 14, 68, -7, 3, 81},
			target:   13,
			expected: [2]int{3, 7},
		},
		{
			nums:     []int{-492, 23, 18, 372, 14, 68, -7, 3, 81},
			target:   -120,
			expected: [2]int{0, 3},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			a, b := twoSum(test.nums, test.target)
			if a+b != test.expected[0]+test.expected[1] {
				t.Errorf("Test failed, expected %v for input %v but got %v\n", test.expected, test.nums, [2]int{a, b})
			}
		},
		)
	}
}
