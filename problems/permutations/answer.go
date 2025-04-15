package main

import (
	"fmt"
	"slices"
)

func permute(nums []int) [][]int {
	return topDownBacktrack(nums)
}

func topDownBacktrack(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	path := make([]int, 0, n)
	used := make([]bool, n)

	var dfs func()
	dfs = func() {
		if len(path) == n {
			res = append(res, slices.Clone(path))
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()

	return res
}

func bottomUp(nums []int) [][]int {
	n := len(nums)
	res := [][]int{{}}

	for i := 0; i < n; i++ {
		newPaths := make([][]int, 0)
		for _, path := range res {
			for _, num := range nums {
				if slices.Contains(path, num) {
					continue
				}
				newPath := slices.Clone(path)
				newPath = append(newPath, num)
				newPaths = append(newPaths, newPath)
			}
		}
		res = newPaths
	}

	return res
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

	solutions := []struct {
		name string
		fn   func([]int) [][]int
	}{
		{
			name: "Default Solution",
			fn:   permute,
		},
		{
			name: "Top-down with backtracking",
			fn:   topDownBacktrack,
		},
		{
			name: "Bottom-up",
			fn:   bottomUp,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if !slices.EqualFunc(actual, tc.expected, func(x, y []int) bool {
				return slices.Equal(x, y)
			}) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
