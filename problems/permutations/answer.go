package main

import (
	"fmt"
	"slices"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, len(nums))
	dfs(nums, make([]bool, len(nums)), &path, &res)
	return res
}

func dfs(nums []int, used []bool, path *[]int, res *[][]int) {
	if len(*path) == len(nums) {
		*res = append(*res, slices.Clone(*path))
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}

		used[i] = true
		*path = append(*path, nums[i])

		dfs(nums, used, path, res)

		*path = (*path)[:len(*path)-1]
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
