package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	n := len(s)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return helper(s, 0, n-1, memo)
}

func helper(s string, l, r int, memo [][]int) int {
	if l > r {
		return 0
	}

	if l == r {
		return 1
	}

	if memo[l][r] != -1 {
		return memo[l][r]
	}

	res := 0
	if s[l] == s[r] {
		res = 2 + helper(s, l+1, r-1, memo)
	} else {
		res = max(helper(s, l+1, r, memo), helper(s, l, r-1, memo))
	}

	memo[l][r] = res
	return res
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "bbbab",
			expected: 4,
		},
		{
			input:    "cbbd",
			expected: 2,
		},
		{
			input:    "abbcccbbbcaaccbababcbcabca",
			expected: 18,
		},
	}

	for _, tc := range testCases {
		actual := longestPalindromeSubseq(tc.input)
		if actual != tc.expected {
			fmt.Printf("longestPalindromeSubseq(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
