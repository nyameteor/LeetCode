package main

import (
	"fmt"
	"slices"
)

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	res := []int{}
	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
		},
		{
			input:    []int{1, 1},
			expected: []int{2},
		},
	}

	for _, tc := range testCases {
		input := slices.Clone(tc.input)
		actual := findDisappearedNumbers(input)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("findDisappearedNumbers(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
