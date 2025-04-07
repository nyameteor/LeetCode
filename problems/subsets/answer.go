package main

import (
	"fmt"
	"slices"
	"sort"
)

func subsets(nums []int) [][]int {
	return iterative(nums)
}

func iterative(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums {
		for _, set := range res {
			newSet := append(slices.Clone(set), num)
			res = append(res, newSet)
		}
	}

	return res
}

func recursive(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	path := make([]int, 0, n)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = append(res, slices.Clone(path))
			return
		}

		dfs(i + 1)

		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}

	dfs(0)

	return res
}

func subsetsEqual(a, b [][]int) bool {
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
		input    []int
		expected [][]int
	}{
		{
			input:    []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			input:    []int{0},
			expected: [][]int{{}, {0}},
		},
	}

	solutions := []struct {
		name string
		fn   func(nums []int) [][]int
	}{
		{
			name: "Default Solution",
			fn:   subsets,
		},
		{
			name: "Recursive with Backtracking",
			fn:   recursive,
		},
		{
			name: "Iterative",
			fn:   iterative,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if !subsetsEqual(actual, tc.expected) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
