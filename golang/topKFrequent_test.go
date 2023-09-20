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
			list:     []int{1, 2},
			k:        2,
			expected: []int{1, 2},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			out := topKFrequent(test.list, test.k)
			if len(out) != len(test.expected) {
				t.Errorf("Expected slice of length %v, got: %v", len(test.expected), len(out))
			}
			sort.Ints(out)
			sort.Ints(test.expected)
			if !reflect.DeepEqual(out, test.expected) {
				t.Errorf("Expected %v, got: %v", test.expected, out)

			}
		})
	}
}
