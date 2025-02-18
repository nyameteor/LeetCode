package main

import (
	"fmt"
	"slices"
)

func permute(nums []int) [][]int {
	var res [][]int
	dfs(nums, make([]bool, len(nums)), make([]int, 0), &res)
	return res
}

func dfs(nums []int, used []bool, path []int, res *[][]int) {
	if len(path) >= len(nums) {
		*res = append(*res, slices.Clone(path))
	}

	for i := range nums {
		if used[i] {
			continue
		}

		used[i] = true
		dfs(nums, used, append(path, nums[i]), res)
		used[i] = false
	}
}

func main() {
	testCases := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			input:    []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			input:    []int{1},
			expected: [][]int{{1}},
		},
	}

	for _, tc := range testCases {
		actual := permute(tc.input)
		if !slices.EqualFunc(actual, tc.expected, func(x, y []int) bool {
			return slices.Equal(x, y)
		}) {
			fmt.Printf("permute(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
