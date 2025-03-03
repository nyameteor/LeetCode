package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}

	return r
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			input:    []int{1, 2, 1, 3, 5, 6, 4},
			expected: 5,
		},
		{
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		actual := findPeakElement(tc.input)
		if actual != tc.expected {
			fmt.Printf("findPeakElement(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
