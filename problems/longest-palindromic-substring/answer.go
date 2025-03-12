package main

import "fmt"

func longestPalindrome(s string) string {
	var res string
	n := len(s)

	helper := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			if r+1-l > len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}
	}

	for i := range n {
		helper(i, i)
		helper(i, i+1)
	}

	return res
}

func main() {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "ccc",
			expected: "ccc",
		},
		{
			input:    "a",
			expected: "a",
		},
		{
			input:    "abbcccbbbcaaccbababcbcabca",
			expected: "bbcccbb",
		},
	}

	for _, tc := range testCases {
		actual := longestPalindrome(tc.input)
		if actual != tc.expected {
			fmt.Printf("longestPalindrome(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
