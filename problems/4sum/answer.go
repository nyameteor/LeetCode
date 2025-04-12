package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, 4, target)
}

func kSum(nums []int, k, target int) [][]int {
	res := [][]int{}

	if len(nums) == 0 {
		return res
	}

	if nums[0] > target/k || target/k > nums[len(nums)-1] {
		return res
	}

	if k == 2 {
		return twoSum(nums, target)
	}

	for i := 0; i < len(nums); i++ {
		// Skip duplicates
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for _, subset := range kSum(nums[i+1:], k-1, target-nums[i]) {
			res = append(res, append([]int{nums[i]}, subset...))
		}
	}

	return res
}

func twoSum(nums []int, target int) [][]int {
	res := [][]int{}
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]+nums[r] < target || (l > 0 && nums[l] == nums[l-1]) {
			l++
		} else if nums[l]+nums[r] > target || (r < len(nums)-1 && nums[r] == nums[r+1]) {
			r--
		} else {
			res = append(res, []int{nums[l], nums[r]})
			l++
			r--
		}
	}

	return res
}

func fourSumDFS(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSumDFS(nums, 4, target)
}

// DFS version for testing and correctness checks, not optimized for performance.
func kSumDFS(nums []int, k, target int) [][]int {
	res := [][]int{}
	subset := make([]int, 0, k)

	var dfs func(start, remain int)
	dfs = func(start, remain int) {
		if len(subset) == k {
			if remain == 0 {
				res = append(res, slices.Clone(subset))
			}
			return
		}

		for i := start; i < len(nums); i++ {
			// Skip duplicates
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			subset = append(subset, nums[i])
			dfs(i+1, remain-nums[i])
			subset = subset[:len(subset)-1]
		}
	}

	dfs(0, target)

	return res
}

func main() {
	type Input struct {
		nums   []int
		target int
	}

	testCases := []struct {
		input    *Input
		expected [][]int
	}{
		{
			input: &Input{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			input: &Input{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			expected: [][]int{{2, 2, 2, 2}},
		},
	}

	solutions := []struct {
		name string
		fn   func(nums []int, target int) [][]int
	}{
		{
			name: "Default Solution",
			fn:   fourSum,
		},
		{
			name: "DFS Solution",
			fn:   fourSumDFS,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input.nums, tc.input.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
