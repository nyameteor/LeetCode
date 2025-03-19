package main

import (
	"fmt"
	"maps"
)

func generateParenthesis(n int) []string {
	res := []string{}
	path := []rune{}
	helper(n, 0, 0, &path, &res)
	return res
}

func helper(n int, open int, close int, path *[]rune, res *[]string) {
	if open == n && close == n {
		*res = append(*res, string(*path))
		return
	}

	if open < n {
		*path = append(*path, '(')
		helper(n, open+1, close, path, res)
		*path = (*path)[:len(*path)-1]
	}

	if close < open {
		*path = append(*path, ')')
		helper(n, open, close+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	testCases := []struct {
		input    int
		expected []string
	}{
		{
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			input:    1,
			expected: []string{"()"},
		},
	}

	setsEqual := func(s1, s2 []string) bool {
		m1, m2 := make(map[string]struct{}), make(map[string]struct{})
		for _, e := range s1 {
			m1[e] = struct{}{}
		}
		for _, s := range s2 {
			m2[s] = struct{}{}
		}
		return maps.Equal(m1, m2)
	}

	for _, tc := range testCases {
		actual := generateParenthesis(tc.input)
		if !setsEqual(actual, tc.expected) {
			fmt.Printf("generateParenthesis(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
