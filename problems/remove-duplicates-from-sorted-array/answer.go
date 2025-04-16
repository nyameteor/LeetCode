package main

import (
	"fmt"
	"slices"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := 0
	r := 1
	for r < len(nums) {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
		r++
	}
	return l + 1
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		input := slices.Clone(tc.input)
		k := removeDuplicates(input)
		actual := input[:k]
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("removeDuplicates(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
