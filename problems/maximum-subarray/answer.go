package main

import (
	"fmt"
	"math"
	"slices"
)

func maxSubArray(nums []int) int {
	return kadane(nums)
}

func kadane(nums []int) int {
	bestSum, curSum := math.MinInt32, 0
	for _, x := range nums {
		curSum = max(x, curSum+x)
		bestSum = max(bestSum, curSum)
	}
	return bestSum
}

func topDownDp(nums []int) int {
	n := len(nums)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	// dfs(i) returns max subarray sum that starts at index i (or later)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return math.MinInt32
		}

		if memo[i] != -1 {
			return memo[i]
		}

		memo[i] = max(nums[i], nums[i]+dfs(i+1))
		return memo[i]
	}

	maxSum := math.MinInt32
	for i := 0; i < n; i++ {
		maxSum = max(maxSum, dfs(i))
	}

	return maxSum
}

func bottomUpDp(nums []int) int {
	n := len(nums)

	// dp[i] stores maximum subarray sum that ends at index i
	dp := slices.Clone(nums)

	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	return slices.Max(dp)
}

func topDownDpWithFlag(nums []int) int {
	n := len(nums)

	type Key struct {
		i         int
		hasPicked bool
	}
	memo := make(map[Key]int)

	var dfs func(i int, hasPicked bool) int
	dfs = func(i int, hasPicked bool) int {
		if i >= n {
			if hasPicked {
				return 0
			}
			// invalid, no element picked
			return math.MinInt32
		}

		key := Key{i, hasPicked}
		if val, ok := memo[key]; ok {
			return val
		}

		if hasPicked {
			// Option 1: Stop picking
			// Option 2: Keep picking
			memo[key] = max(0, nums[i]+dfs(i+1, true))
			return memo[key]
		}

		// Option 1: Skip this element, still haven't picked any
		// Option 2: Start picking from here
		memo[key] = max(dfs(i+1, false), nums[i]+dfs(i+1, true))
		return memo[key]
	}

	return dfs(0, false)
}

func divideAndConquer(nums []int) int {
	var maxRangeSum func(left, right int) int
	maxRangeSum = func(left, right int) int {
		if left > right {
			return math.MinInt32
		}

		mid := left + (right-left)/2

		leftMax := 0
		for i, curSum := mid-1, 0; i >= left; i-- {
			curSum += nums[i]
			leftMax = max(leftMax, curSum)
		}

		rightMax := 0
		for i, curSum := mid+1, 0; i <= right; i++ {
			curSum += nums[i]
			rightMax = max(rightMax, curSum)
		}

		crossSum := nums[mid] + leftMax + rightMax

		leftSum := maxRangeSum(left, mid-1)
		rightSum := maxRangeSum(mid+1, right)

		return max(leftSum, rightSum, crossSum)
	}

	n := len(nums)
	return maxRangeSum(0, n-1)
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			input:    []int{1},
			expected: 1,
		},
		{
			input:    []int{5, 4, -1, 7, 8},
			expected: 23,
		},
	}

	solutions := []struct {
		name string
		fn   func([]int) int
	}{
		{
			name: "Default",
			fn:   maxSubArray,
		},
		{
			name: "Kadane's Algorithm",
			fn:   kadane,
		},
		{
			name: "Top-down DP",
			fn:   topDownDp,
		},
		{
			name: "Bottom-up DP",
			fn:   bottomUpDp,
		},
		{
			name: "Top-down DP (With Flag)",
			fn:   topDownDpWithFlag,
		},
		{
			name: "Divide and Conquer",
			fn:   divideAndConquer,
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
