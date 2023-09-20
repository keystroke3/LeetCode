package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := groupAnagrams(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
			}
		})
	}
}
