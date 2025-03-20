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

func longestPalindromeDp1(s string) string {
	var res string
	n := len(s)

	memo := make(map[[2]int]bool)

	var helper func(l, r int) bool
	helper = func(l, r int) bool {
		key := [2]int{l, r}
		if val, ok := memo[key]; ok {
			return val
		}

		if l >= r {
			memo[key] = true
			return memo[key]
		}

		if s[l] == s[r] && helper(l+1, r-1) {
			memo[key] = true
			return memo[key]
		}

		memo[key] = false
		return memo[key]
	}

	// Iterate from longest to shortest substrings to maximize cache hits.
	for length := n; length > 0; length-- {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			if helper(l, r) {
				res = s[l : r+1]
				return res
			}
		}
	}

	return res
}

func longestPalindromeDp2(s string) string {
	var res string
	n := len(s)

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1

			if length == 1 {
				dp[l][r] = true
			} else if length == 2 {
				dp[l][r] = (s[l] == s[r])
			} else {
				dp[l][r] = (s[l] == s[r]) && dp[l+1][r-1]
			}

			if dp[l][r] && r-l+1 > len(res) {
				res = s[l : r+1]
			}
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    string
		expected map[string]struct{}
	}{
		{
			input:    "babad",
			expected: map[string]struct{}{"bab": {}, "aba": {}},
		},
		{
			input:    "cbbd",
			expected: map[string]struct{}{"bb": {}},
		},
		{
			input:    "ccc",
			expected: map[string]struct{}{"ccc": {}},
		},
		{
			input:    "a",
			expected: map[string]struct{}{"a": {}},
		},
		{
			input:    "abbcccbbbcaaccbababcbcabca",
			expected: map[string]struct{}{"bbcccbb": {}, "cbababc": {}},
		},
	}

	type Solution struct {
		name string
		fn   func(s string) string
	}

	solutions := []Solution{
		{
			name: "Two Pointers",
			fn:   longestPalindrome,
		},
		{
			name: "Dynamic Programming (Top-Down)",
			fn:   longestPalindromeDp1,
		},
		{
			name: "Dynamic Programming (Bottom-Up)",
			fn:   longestPalindromeDp2,
		},
	}

	for _, tc := range testCases {
		for _, solution := range solutions {
			actual := solution.fn(tc.input)
			if _, ok := tc.expected[actual]; !ok {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
