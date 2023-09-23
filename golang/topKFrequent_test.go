package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopFrequent(t *testing.T) {
	tests := []struct {
		list     []int
		k        int
		expected []int
	}{
		{
			list:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			list:     []int{4, 4, 4, 6, 6, 7},
			k:        3,
			expected: []int{4, 6, 7},
		},
		{
			list:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			k:        10,
			expected: []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11},
		},
		{
			list:     []int{99, 100},
			k:        2,
			expected: []int{99, 100},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			out := topKFrequent(test.list, test.k)
			if len(out) != len(test.expected) {
				t.Errorf("Expected slice of length %v, got: %v", len(test.expected), len(out))
			}
			var outCopy, expCopy []int
			copy(outCopy, out)
			copy(expCopy, test.expected)
			sort.Ints(outCopy)
			sort.Ints(expCopy)
			if !reflect.DeepEqual(out, test.expected) {
				t.Errorf("Expected %v, got: %v", test.expected, out)

			}
		})
	}
}
