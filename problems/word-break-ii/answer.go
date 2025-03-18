package main

import (
	"fmt"
	"maps"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	memo := make(map[string][]string)
	return helper(s, wordDict, memo)
}

func helper(s string, wordDict []string, memo map[string][]string) []string {
	if res, ok := memo[s]; ok {
		return res
	}

	if len(s) == 0 {
		return []string{""}
	}

	res := make([]string, 0)
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			subSentences := helper(s[len(word):], wordDict, memo)
			for _, sub := range subSentences {
				if sub == "" {
					res = append(res, word)
				} else {
					res = append(res, word+" "+sub)
				}
			}
		}
	}

	memo[s] = res
	return res
}

func main() {
	type Input struct {
		s        string
		wordDict []string
	}

	testCases := []struct {
		input    *Input
		expected []string
	}{
		{
			input: &Input{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			expected: []string{"cats and dog", "cat sand dog"},
		},
		{
			input: &Input{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			expected: []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			input: &Input{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			expected: []string{},
		},
	}

	setsEqual := func(s1, s2 []string) bool {
		m1, m2 := make(map[string]struct{}), make(map[string]struct{})
		for _, e := range s1 {
			m1[e] = struct{}{}
		}
		for _, e := range s2 {
			m2[e] = struct{}{}
		}
		return maps.Equal(m1, m2)
	}

	for _, tc := range testCases {
		actual := wordBreak(tc.input.s, tc.input.wordDict)
		if !setsEqual(actual, tc.expected) {
			fmt.Printf("wordBreak(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
