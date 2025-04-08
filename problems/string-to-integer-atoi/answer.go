package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {
	runes := []rune(s)

	n := len(runes)
	i := 0

	for ; i < n && unicode.IsSpace(runes[i]); i++ {
	}

	sign := 1
	if i < n && runes[i] == '-' {
		sign = -1
		i++
	} else if i < n && runes[i] == '+' {
		i++
	}

	for ; i < n && runes[i] == '0'; i++ {
	}

	var res int64 = 0
	for ; i < n && unicode.IsDigit(runes[i]); i++ {
		res = res*10 + int64(runes[i]-'0')

		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && -res < math.MinInt32 {
			return math.MinInt32
		}
	}

	return sign * int(res)
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "42",
			expected: 42,
		},
		{
			input:    " -042",
			expected: -42,
		},
		{
			input:    "1337c0d3",
			expected: 1337,
		},
		{
			input:    "0-1",
			expected: 0,
		},
		{
			input:    "words and 987",
			expected: 0,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "21474836460",
			expected: math.MaxInt32,
		},
	}

	for _, tc := range testCases {
		actual := myAtoi(tc.input)
		if actual != tc.expected {
			fmt.Printf("myAtoi(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
