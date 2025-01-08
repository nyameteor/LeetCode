package main

import "fmt"

func countPalindromicSubsequence(s string) int {
	n := len(s)

	res := 0

	for x := 'a'; x <= 'z'; x++ {
		first, last := -1, -1

		for i := 0; i < n; i++ {
			if rune(s[i]) == x {
				if first == -1 {
					first = i
				}
				last = i
			}
		}

		if first == -1 || last == -1 || first == last {
			continue
		}

		seen := map[rune]struct{}{}
		for i := first + 1; i < last; i++ {
			seen[rune(s[i])] = struct{}{}
		}

		res += len(seen)
	}

	return res
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "aabca",
			expected: 3,
		},
		{
			input:    "adc",
			expected: 0,
		},
		{
			input:    "bbcbaba",
			expected: 4,
		},
	}

	for _, tc := range testCases {
		actual := countPalindromicSubsequence(tc.input)
		if actual != tc.expected {
			fmt.Printf("countPalindromicSubsequence(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
