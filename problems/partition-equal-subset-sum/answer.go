package main

import "fmt"

func canPartition(nums []int) bool {
	return bottomUpDp(nums)
}

func topDownDp(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2

	n := len(nums)
	memo := make(map[[2]int]bool)

	var dfs func(i, sum int) bool
	dfs = func(i, sum int) bool {
		if sum == target {
			return true
		}

		if i >= n || sum > target {
			return false
		}

		key := [2]int{i, sum}
		if val, ok := memo[key]; ok {
			return val
		}

		memo[key] = dfs(i+1, sum) || dfs(i+1, sum+nums[i])
		return memo[key]
	}

	return dfs(0, 0)
}

func bottomUpDp(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2

	n := len(nums)
	dp := make([]map[int]bool, n+1)
	for i := range dp {
		dp[i] = make(map[int]bool)
	}
	dp[0][0] = true

	for i := 1; i <= n; i++ {
		num := nums[i-1]
		for sum := range dp[i-1] {
			dp[i][sum] = true
			dp[i][sum+num] = true
		}
	}

	return dp[n][target]
}

func main() {
	testCases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 5, 11, 5},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 5},
			expected: false,
		},
	}

	solutions := []struct {
		name string
		fn   func([]int) bool
	}{
		{
			name: "Default Solution",
			fn:   canPartition,
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
