package main

import (
	"fmt"
	"slices"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && (temperatures[stack[len(stack)-1]] <= temperatures[i]) {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			input:    []int{30, 40, 50, 60},
			expected: []int{1, 1, 1, 0},
		},
		{
			input:    []int{30, 60, 90},
			expected: []int{1, 1, 0},
		},
		{
			input:    []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70},
			expected: []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0},
		},
	}

	for _, tc := range testCases {
		actual := dailyTemperatures(tc.input)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("dailyTemperatures(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
