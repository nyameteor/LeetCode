package main

import (
	"fmt"
)

func rob(nums []int) int {
	return bottomUp(nums)
}

func bottomUp(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	return dp[n-1]
}

func topDown(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		memo[i] = max(dfs(i-1), nums[i]+dfs(i-2))
		return memo[i]
	}

	return dfs(len(nums) - 1)
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			input:    []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			input:    []int{0},
			expected: 0,
		},
	}

	solutions := []struct {
		name string
		fn   func(nums []int) int
	}{
		{
			name: "Default Solution",
			fn:   rob,
		},
		{
			name: "Top-down DP",
			fn:   topDown,
		},
		{
			name: "Bottom-up DP",
			fn:   bottomUp,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			output := solution.fn(tc.input)
			if output != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, output, tc.expected)
			}
		}
	}
}
