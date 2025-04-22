package main

import (
	"fmt"
	"slices"
)

func sortColors(nums []int) {
	mid := 1
	i, j, k := 0, 0, len(nums)-1

	for j <= k {
		if nums[j] < mid {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] > mid {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else {
			j++
		}
	}
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
	}

	for _, tc := range testCases {
		actual := slices.Clone(tc.input)
		sortColors(actual)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("sortColors(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
