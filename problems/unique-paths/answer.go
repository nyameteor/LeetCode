package main

import "fmt"

func uniquePaths(m int, n int) int {
	return bottomUp(m, n)
}

func bottomUp(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		for j := range dp[i] {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func topDown(m int, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}
		if i == m-1 && j == n-1 {
			return 1
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		memo[i][j] = dfs(i+1, j) + dfs(i, j+1)
		return memo[i][j]
	}

	return dfs(0, 0)
}

func main() {
	type Input struct {
		m, n int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				m: 3,
				n: 7,
			},
			expected: 28,
		},
		{
			input: &Input{
				m: 3,
				n: 2,
			},
			expected: 3,
		},
	}

	solutions := []struct {
		name string
		fn   func(m, n int) int
	}{
		{
			name: "Default Solution",
			fn:   uniquePaths,
		},
		{
			name: "Top-down",
			fn:   topDown,
		},
		{
			name: "Bottom-up",
			fn:   bottomUp,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input.m, tc.input.n)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
