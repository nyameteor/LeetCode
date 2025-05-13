package main

import "fmt"

func longestOnes(nums []int, k int) int {
	res := 0

	zeroCount := 0
	start := 0

	for i := range nums {
		if nums[i] == 0 {
			zeroCount++
		}

		// Shrink window if we have more than k zeros
		for zeroCount > k {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}

		res = max(res, i-start+1)
	}

	return res
}

func main() {
	type Input struct {
		nums []int
		k    int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    2,
			},
			expected: 6,
		},
		{
			input: &Input{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    3,
			},
			expected: 10,
		},
		{
			input: &Input{
				nums: []int{0, 0, 1, 1, 1, 0, 0},
				k:    0,
			},
			expected: 3,
		},
		{
			input: &Input{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    0,
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		actual := longestOnes(tc.input.nums, tc.input.k)
		if actual != tc.expected {
			fmt.Printf("longestOnes(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
