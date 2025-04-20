package main

import "fmt"

func majorityElement(nums []int) int {
	res := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			res = num
			count = 1
		} else if num == res {
			count++
		} else {
			count--
		}
	}
	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{3, 2, 3},
			expected: 3,
		},
		{
			input:    []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		actual := majorityElement(tc.input)
		if actual != tc.expected {
			fmt.Printf("majorityElement(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
