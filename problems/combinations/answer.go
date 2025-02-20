package main

import (
	"fmt"
	"slices"
)

func combine(n int, k int) [][]int {
	var res [][]int
	dfs(n, k, 1, make([]int, 0), &res)
	return res
}

func dfs(n int, k int, start int, path []int, res *[][]int) {
	if len(path) >= k {
		*res = append(*res, slices.Clone(path))
		return
	}

	for i := start; i <= n; i++ {
		dfs(n, k, i+1, append(path, i), res)
	}
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

	for _, tc := range testCases {
		actual := combine(tc.input.n, tc.input.k)
		if !slices.EqualFunc(actual, tc.expected, func(x, y []int) bool {
			return slices.Equal(x, y)
		}) {
			fmt.Printf("combine(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
