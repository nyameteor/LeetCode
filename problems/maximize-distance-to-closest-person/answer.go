package main

import (
	"fmt"
)

func maxDistToClosest(seats []int) int {
	n := len(seats)

	res := 0
	last := -1

	for i := range seats {
		if seats[i] == 0 {
			continue
		}
		if last == -1 {
			res = i
		} else {
			res = max(res, (i-last)/2)
		}
		last = i
	}
	res = max(res, n-1-last)

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 0, 0, 0, 1, 0, 1},
			expected: 2,
		},
		{
			input:    []int{1, 0, 0, 0},
			expected: 3,
		},
		{
			input:    []int{0, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := maxDistToClosest(tc.input)
		if actual != tc.expected {
			fmt.Printf("maxDistToClosest(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
