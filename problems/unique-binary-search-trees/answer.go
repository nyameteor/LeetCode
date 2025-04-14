package main

import "fmt"

func numTrees(n int) int {
	memo := make([]int, n+1)

	var countTrees func(size int) int
	countTrees = func(size int) int {
		if size < 0 {
			return 0
		}
		if size == 0 {
			return 1
		}

		if memo[size] != 0 {
			return memo[size]
		}

		res := 0
		for left := 0; left < size; left++ {
			right := size - 1 - left
			res += countTrees(left) * countTrees(right)
		}
		memo[size] = res
		return res
	}

	return countTrees(n)
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    3,
			expected: 5,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    5,
			expected: 42,
		},
	}

	for _, tc := range testCases {
		actual := numTrees(tc.input)
		if actual != tc.expected {
			fmt.Printf("numTrees(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
