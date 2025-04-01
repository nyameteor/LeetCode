package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

func combine(n int, k int) [][]int {
	return topDownBacktrack(n, k)
}

func topDownBacktrack(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, k)

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			res = append(res, slices.Clone(path))
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return res
}

func topDown(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		if len(path) == k {
			res = append(res, slices.Clone(path))
			return
		}

		for i := start; i <= n; i++ {
			dfs(i+1, append(path, i))
		}
	}

	dfs(1, path)
	return res
}

func combinationsEqual(c1, c2 [][]int) bool {
	if len(c1) != len(c2) {
		return false
	}
	for i := range c1 {
		if len(c1[i]) != len(c2[i]) {
			return false
		}
	}

	for i := range c1 {
		sort.Ints(c1[i])
		sort.Ints(c2[i])
	}

	slicesCompare := func(s1, s2 []int) bool {
		for i := range s1 {
			if s1[i] != s2[i] {
				return s1[i] < s2[i]
			}
		}
		return false
	}

	sort.Slice(c1, func(i, j int) bool { return slicesCompare(c1[i], c1[j]) })
	sort.Slice(c2, func(i, j int) bool { return slicesCompare(c2[i], c2[j]) })

	return reflect.DeepEqual(c1, c2)
}

func main() {
	type Input struct {
		n int
		k int
	}

	testCases := []struct {
		input    Input
		expected [][]int
	}{
		{
			input: Input{
				n: 4,
				k: 2,
			},
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			input: Input{
				n: 1,
				k: 1,
			},
			expected: [][]int{{1}},
		},
		{
			input: Input{
				n: 5,
				k: 4,
			},
			expected: [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}},
		},
	}

	solutions := []struct {
		name string
		fn   func(n int, k int) [][]int
	}{
		{
			name: "Default Solution",
			fn:   combine,
		},
		{
			name: "Top-down",
			fn:   topDown,
		},
		{
			name: "Top-down with backtracking",
			fn:   topDownBacktrack,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input.n, tc.input.k)
			if !combinationsEqual(actual, tc.expected) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
