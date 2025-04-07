package main

import "fmt"

func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1

	for l < r {
		res = max(res, (r-l)*min(height[l], height[r]))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			input:    []int{1, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := maxArea(tc.input)
		if actual != tc.expected {
			fmt.Printf("maxArea(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
