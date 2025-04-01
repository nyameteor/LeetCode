package main

import "fmt"

func change(amount int, coins []int) int {
	return bottomUpDp(amount, coins)
}

func bottomUpDp(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func topDownDp(amount int, coins []int) int {
	memo := make([][]int, amount+1)
	for i := range memo {
		memo[i] = make([]int, len(coins))
		for j := range coins {
			memo[i][j] = -1
		}
	}

	var dfs func(remain int, i int) int
	dfs = func(remain int, i int) int {
		if remain < 0 {
			return 0
		}
		if remain == 0 {
			return 1
		}

		if memo[remain][i] != -1 {
			return memo[remain][i]
		}

		res := 0
		for j := i; j < len(coins); j++ {
			res += dfs(remain-coins[j], j)
		}

		memo[remain][i] = res
		return res
	}

	return dfs(amount, 0)
}

func main() {
	type Input struct {
		amount int
		coins  []int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			expected: 4,
		},
		{
			input: &Input{
				amount: 3,
				coins:  []int{2},
			},
			expected: 0,
		},
		{
			input: &Input{
				amount: 10,
				coins:  []int{10},
			},
			expected: 1,
		},
		{
			input: &Input{
				amount: 500,
				coins:  []int{3, 5, 7, 8, 9, 10, 11},
			},
			expected: 35502874,
		},
	}

	solutions := []struct {
		name string
		fn   func(amount int, coins []int) int
	}{
		{
			name: "Default Solution",
			fn:   change,
		},
		{
			name: "Bottom-up DP",
			fn:   bottomUpDp,
		},
		{
			name: "Top-down DP",
			fn:   topDownDp,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input.amount, tc.input.coins)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
