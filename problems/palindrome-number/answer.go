package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == original
}

func main() {
	testCases := []struct {
		input    int
		expected bool
	}{
		{
			input:    121,
			expected: true,
		},
		{
			input:    -121,
			expected: false,
		},
		{
			input:    10,
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := isPalindrome(tc.input)
		if actual != tc.expected {
			fmt.Printf("isPalindrome(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
