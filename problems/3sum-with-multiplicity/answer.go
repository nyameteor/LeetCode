package main

import (
	"fmt"
)

func threeSumMulti(arr []int, target int) int {
	const mod int = 1e9 + 7

	res := 0
	freq := make(map[int]int)

	for i := range arr {
		res = res + freq[target-arr[i]]
		res = res % mod

		for j := range i {
			freq[arr[i]+arr[j]]++
		}
	}

	return res
}

func main() {
	type Input struct {
		arr    []int
		target int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				arr:    []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
				target: 8,
			},
			expected: 20,
		},
		{
			input: &Input{
				arr:    []int{1, 1, 2, 2, 2, 2},
				target: 5,
			},
			expected: 12,
		},
		{
			input: &Input{
				arr:    []int{2, 1, 3},
				target: 6,
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := threeSumMulti(tc.input.arr, tc.input.target)
		if actual != tc.expected {
			fmt.Printf("threeSumMulti(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
