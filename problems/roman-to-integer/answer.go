package main

import "fmt"

var romanValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	for i := range s {
		cur := romanValues[s[i]]
		if i+1 < len(s) && cur < romanValues[s[i+1]] {
			res -= cur
		} else {
			res += cur
		}
	}
	return res
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "III",
			expected: 3,
		},
		{
			input:    "LVIII",
			expected: 58,
		},
		{
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tc := range testCases {
		actual := romanToInt(tc.input)
		if actual != tc.expected {
			fmt.Printf("romanToInt(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
