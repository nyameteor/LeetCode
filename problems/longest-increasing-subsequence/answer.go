package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	return bottomUpDp(nums)
}

func bottomUpDp(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	res := 0
	for i := range n {
		for j := range i {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func topDownDp(nums []int) int {
	memo := make(map[int]int)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}

		if v, ok := memo[i]; ok {
			return v
		}

		res := 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				res = max(res, 1+dfs(j))
			}
		}
		memo[i] = res
		return res
	}

	res := 0
	for i := range nums {
		res = max(res, dfs(i))
	}
	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			input:    []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			input:    []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			input:    []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			expected: 6,
		},
	}

	solutions := []struct {
		name string
		fn   func(nums []int) int
	}{
		{
			name: "Default Solution",
			fn:   lengthOfLIS,
		},
		{
			name: "Top-down DP",
			fn:   topDownDp,
		},
		{
			name: "Bottom-up DP",
			fn:   bottomUpDp,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
