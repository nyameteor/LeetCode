package main

import "fmt"

func climbStairs(n int) int {
	return bottomUpSpaceOptimized(n)
}

func topDown(n int) int {
	memo := make([]int, n+1)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return 1
		}
		if memo[i] != 0 {
			return memo[i]
		}

		memo[i] = dfs(i-1) + dfs(i-2)
		return memo[i]
	}

	return dfs(n)
}

func bottomUp(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func bottomUpSpaceOptimized(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	pre, cur := 1, 1
	for i := 2; i <= n; i++ {
		pre, cur = cur, pre+cur
	}

	return cur
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    5,
			expected: 8,
		},
	}

	solutions := []struct {
		name string
		fn   func(int) int
	}{
		{
			name: "Default Solution",
			fn:   climbStairs,
		},
		{
			name: "Top-down DP",
			fn:   topDown,
		},
		{
			name: "Bottom-up DP",
			fn:   bottomUp,
		},
		{
			name: "Bottom-up DP (Space Optimized)",
			fn:   bottomUpSpaceOptimized,
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
