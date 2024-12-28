package main

import (
	"fmt"
	"math"
	"slices"
)

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)

	subarraySums := computeSubarraySums(nums, k)
	bestLeftSums, bestLeftIdxes := computeBestLeft(subarraySums)
	bestRightSums, bestRightIdxes := computeBestRight(subarraySums)

	bestSum := math.MinInt32
	bestPoses := []int{-1, -1, -1}

	for i := k; i <= n-2*k; i++ {
		m := i
		l := i - k
		r := i + k

		curSum := subarraySums[m] + bestLeftSums[l] + bestRightSums[r]
		if bestSum < curSum {
			bestSum = curSum
			bestPoses = []int{bestLeftIdxes[l], m, bestRightIdxes[r]}
		}
	}

	return bestPoses
}

func computeSubarraySums(nums []int, k int) []int {
	n := len(nums)
	if n < k {
		return nil
	}

	sums := make([]int, n-k+1)

	curSum := 0
	for i := 0; i < k; i++ {
		curSum += nums[i]
	}
	sums[0] = curSum

	for i := 1; i <= n-k; i++ {
		curSum = curSum - nums[i-1] + nums[i+k-1]
		sums[i] = curSum
	}

	return sums
}

func computeBestLeft(subarraySums []int) ([]int, []int) {
	n := len(subarraySums)
	if n == 0 {
		return nil, nil
	}

	bestSums := make([]int, n)
	bestIdxes := make([]int, n)

	bestSum := math.MinInt32
	bestIdx := -1

	for i := 0; i < n; i++ {
		curSum := subarraySums[i]
		if curSum > bestSum {
			bestSum = curSum
			bestIdx = i
		}
		bestSums[i] = bestSum
		bestIdxes[i] = bestIdx
	}

	return bestSums, bestIdxes
}

func computeBestRight(subarraySums []int) ([]int, []int) {
	n := len(subarraySums)
	if n == 0 {
		return nil, nil
	}

	bestSums := make([]int, n)
	bestIdxes := make([]int, n)

	bestSum := math.MinInt32
	bestIdx := -1

	for i := n - 1; i >= 0; i-- {
		curSum := subarraySums[i]
		if curSum >= bestSum {
			bestSum = curSum
			bestIdx = i
		}
		bestSums[i] = bestSum
		bestIdxes[i] = bestIdx
	}

	return bestSums, bestIdxes
}

func main() {
	type Input struct {
		nums []int
		k    int
	}
	type Output struct {
		positions []int
	}

	testCases := []struct {
		input    Input
		expected Output
	}{
		{
			Input{
				nums: []int{1, 2, 1, 2, 6, 7, 5, 1},
				k:    2,
			},
			Output{
				positions: []int{0, 3, 5},
			},
		},
		{
			Input{
				nums: []int{1, 2, 1, 2, 1, 2, 1, 2, 1},
				k:    2,
			},
			Output{
				positions: []int{0, 2, 4},
			},
		},
	}

	for _, tc := range testCases {
		actual := maxSumOfThreeSubarrays(tc.input.nums, tc.input.k)
		if !slices.Equal(tc.expected.positions, actual) {
			fmt.Printf("maxSumOfThreeSubarrays(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
