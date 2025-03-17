package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)
	return helper(s, wordDict, memo)
}

func helper(s string, wordDict []string, memo map[string]bool) bool {
	if val, ok := memo[s]; ok {
		return val
	}

	if len(s) == 0 {
		return true
	}

	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			if helper(s[len(word):], wordDict, memo) {
				memo[s] = true
				return memo[s]
			}
		}
	}

	memo[s] = false
	return memo[s]
}

func main() {
	type Input struct {
		s        string
		wordDict []string
	}

	testCases := []struct {
		input    *Input
		expected bool
	}{
		{
			input: &Input{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			expected: true,
		},
		{
			input: &Input{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			expected: true,
		},
		{
			input: &Input{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			expected: false,
		},
		{
			input: &Input{
				s:        "ccbb",
				wordDict: []string{"bc", "cb"},
			},
			expected: false,
		},
		{
			input: &Input{
				s:        "cars",
				wordDict: []string{"car", "ca", "rs"},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		actual := wordBreak(tc.input.s, tc.input.wordDict)
		if actual != tc.expected {
			fmt.Printf("wordBreak(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
