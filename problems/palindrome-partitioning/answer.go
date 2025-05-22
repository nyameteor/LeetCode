package main

import (
	"fmt"
	"slices"
	"sort"
)

func partition(s string) [][]string {
	// check palindrome with memoization
	memo := make(map[[2]int]bool)
	palindrome := func(l, r int) bool {
		key := [2]int{l, r}
		if val, ok := memo[key]; ok {
			return val
		}
		memo[key] = isPalindrome(s, l, r)
		return memo[key]
	}

	n := len(s)
	res := [][]string{}
	path := []string{}

	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			res = append(res, slices.Clone(path))
			return
		}

		for i := start; i < n; i++ {
			if !palindrome(start, i) {
				continue
			}

			path = append(path, s[start:i+1])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return res
}

func isPalindrome(s string, l, r int) bool {
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}

func subsetsEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		slices.Sort(a[i])
	}
	for i := range b {
		slices.Sort(b[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return slices.Compare(a[i], a[j]) < 0
	})
	sort.Slice(b, func(i, j int) bool {
		return slices.Compare(b[i], b[j]) < 0
	})

	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		input    string
		expected [][]string
	}{
		{
			input:    "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			input:    "a",
			expected: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		actual := partition(tc.input)
		if !subsetsEqual(actual, tc.expected) {
			fmt.Printf("partition(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
