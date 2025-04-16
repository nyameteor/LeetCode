package main

import (
	"fmt"
	"slices"
)

func nextPermutation(nums []int) {
	n := len(nums)

	k := n - 2
	for k >= 0 && nums[k] >= nums[k+1] {
		k--
	}

	if k < 0 {
		reverse(nums, 0, n-1)
		return
	}

	l := n - 1
	for l > k && nums[l] <= nums[k] {
		l--
	}

	nums[k], nums[l] = nums[l], nums[k]

	reverse(nums, k+1, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			input:    []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
	}

	for _, tc := range testCases {
		actual := slices.Clone(tc.input)
		nextPermutation(actual)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("nextPermutation(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
