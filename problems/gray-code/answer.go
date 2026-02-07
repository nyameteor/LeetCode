package main

import (
	"fmt"
	"slices"
)

func grayCode(n int) []int {
	res := make([]int, 1, 1<<n)
	res[0] = 0

	for i := 1; i <= n; i++ {
		bit := 1 << (i - 1)
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]|bit)
		}
	}
	return res
}

func main() {
	testCases := []struct {
		input    int
		expected []int
	}{
		{
			input:    1,
			expected: []int{0, 1},
		},
		{
			input:    2,
			expected: []int{0, 1, 3, 2},
		},
		{
			input:    3,
			expected: []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
	}

	for _, tc := range testCases {
		actual := grayCode(tc.input)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("grayCode(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
