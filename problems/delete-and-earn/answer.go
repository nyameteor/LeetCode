package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	vals := []int{}
	for num := range freq {
		vals = append(vals, num)
	}
	sort.Ints(vals)

	memo := make([]int, len(vals))
	for i := range memo {
		memo[i] = -1
	}

	return helper(freq, vals, memo, len(vals)-1)
}

func helper(freq map[int]int, vals []int, memo []int, i int) int {
	if i < 0 {
		return 0
	}

	if memo[i] != -1 {
		return memo[i]
	}

	earn := vals[i] * freq[vals[i]]
	if i-1 >= 0 && vals[i]-1 == vals[i-1] {
		memo[i] = max(helper(freq, vals, memo, i-1), earn+helper(freq, vals, memo, i-2))
	} else {
		memo[i] = earn + helper(freq, vals, memo, i-1)
	}

	return memo[i]
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{3, 4, 2},
			expected: 6,
		},
		{
			input:    []int{2, 2, 3, 3, 3, 4},
			expected: 9,
		},
		{
			input:    []int{1, 6, 3, 3, 8, 4, 8, 10, 1, 3},
			expected: 43,
		},
		{
			input:    []int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10},
			expected: 37,
		},
		{
			input:    []int{1, 1, 1, 2, 4, 5, 5, 5, 6},
			expected: 18,
		},
	}

	for _, tc := range testCases {
		actual := deleteAndEarn(tc.input)
		if actual != tc.expected {
			fmt.Printf("deleteAndEarn(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
