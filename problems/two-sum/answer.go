package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)
	for i, num := range nums {
		if j, ok := numIndices[target-num]; ok {
			return []int{j, i}
		}
		numIndices[num] = i
	}
	return nil
}

func main() {
	type Input struct {
		nums   []int
		target int
	}

	testCases := []struct {
		input    *Input
		expected []int
	}{
		{
			input: &Input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			input: &Input{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			input: &Input{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		actual := twoSum(tc.input.nums, tc.input.target)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("twoSum(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
