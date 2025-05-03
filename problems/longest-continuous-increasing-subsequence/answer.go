package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	res := 0

	for l, r := 0, 0; r < len(nums); r++ {
		if r > l && nums[r] <= nums[r-1] {
			l = r
		}
		res = max(res, r-l+1)
	}

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 3, 5, 4, 7},
			expected: 3,
		},
		{
			input:    []int{2, 2, 2, 2, 2},
			expected: 1,
		},
		{
			input:    []int{1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := findLengthOfLCIS(tc.input)
		if actual != tc.expected {
			fmt.Printf("findLengthOfLCIS(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
