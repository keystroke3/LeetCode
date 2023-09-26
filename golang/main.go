package main

func main() {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "race car",
			expected: true,
		},
		// {
		// 	input:    "A man, a plan, a canal: Panama",
		// 	expected: true,
		// },
		// {
		// 	input:    "  ",
		// 	expected: false,
		// },
		// {
		// 	input:    " ",
		// 	expected: false,
		// },
		// {
		// 	input:    "aa",
		// 	expected: true,
		// },
		// {
		// 	input:    " man, a plan, a canal: Panama",
		// 	expected: false,
		// },

		// {
		// 	input:    "race a car",
		// 	expected: false,
		// },
	}
	for _, t := range tests {
		isPalindrome(t.input)
	}
}
