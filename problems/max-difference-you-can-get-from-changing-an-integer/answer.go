package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	s := strconv.Itoa(num)

	maxStr := s
	for i := range s {
		if s[i] != '9' {
			maxStr = strings.ReplaceAll(s, string(s[i]), "9")
			break
		}
	}

	minStr := s
	if s[0] != '1' {
		minStr = strings.ReplaceAll(s, string(s[0]), "1")
	} else {
		for i := 1; i < len(s); i++ {
			if s[i] != '0' && s[i] != s[0] {
				minStr = strings.ReplaceAll(s, string(s[i]), "0")
				break
			}
		}
	}

	maxNum, _ := strconv.Atoi(maxStr)
	minNum, _ := strconv.Atoi(minStr)

	return maxNum - minNum
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    555,
			expected: 888,
		},
		{
			input:    9,
			expected: 8,
		},
		{
			input:    1101057,
			expected: 8808050,
		},
	}

	for _, tc := range testCases {
		actual := maxDiff(tc.input)
		if actual != tc.expected {
			fmt.Printf("maxDiff(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
