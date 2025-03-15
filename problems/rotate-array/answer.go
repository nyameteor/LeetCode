package main

import (
	"fmt"
	"slices"
)

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse[T any](s []T, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l, r = l+1, r-1
	}
}

func main() {
	type Input struct {
		nums []int
		k    int
	}

	testCases := []struct {
		input    *Input
		expected []int
	}{
		{
			input: &Input{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			input: &Input{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expected: []int{3, 99, -1, -100},
		},
	}

	for _, tc := range testCases {
		numsCopy := slices.Clone(tc.input.nums)
		rotate(tc.input.nums, tc.input.k)
		actual := tc.input.nums
		tc.input.nums = numsCopy
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("rotate(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
