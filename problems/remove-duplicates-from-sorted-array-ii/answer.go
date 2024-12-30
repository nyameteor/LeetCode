package main

import (
	"fmt"
	"slices"
)

func removeDuplicates(nums []int) int {
	const maxDupCount = 2

	n := len(nums)

	if n <= maxDupCount {
		return n
	}

	dupCount := 1
	l, r := 1, 1

	for ; r < n; r++ {
		if nums[r] == nums[r-1] {
			dupCount++
		} else {
			dupCount = 1
		}

		if dupCount <= maxDupCount {
			nums[l] = nums[r]
			l++
		}
	}

	return l
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			[]int{1, 1, 2, 2, 3},
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for _, tc := range testCases {
		inputCopy := slices.Clone(tc.input)
		k := removeDuplicates(tc.input)
		result := tc.input[:k]
		if !slices.Equal(result, tc.expected) {
			fmt.Printf("Test failed for input: %v, Result: %v, Expected: %v\n", inputCopy, result, tc.expected)
		}
	}
}
