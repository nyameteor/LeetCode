package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 2, 1},
			expected: 1,
		},
		{
			input:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			input:    []int{1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := singleNumber(tc.input)
		if actual != tc.expected {
			fmt.Printf("singleNumber(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
