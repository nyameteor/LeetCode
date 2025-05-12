package main

import "fmt"

func numPairsDivisibleBy60(time []int) int {
	remainders := make([]int, 60)

	for i := range time {
		remainders[time[i]%60]++
	}

	res := 0
	for i := 1; i < 30; i++ {
		res += remainders[i] * remainders[60-i]
	}

	res += remainders[0] * (remainders[0] - 1) / 2
	res += remainders[30] * (remainders[30] - 1) / 2

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{30, 20, 150, 100, 40},
			expected: 3,
		},
		{
			input:    []int{60, 60, 60},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		actual := numPairsDivisibleBy60(tc.input)
		if actual != tc.expected {
			fmt.Printf("numPairsDivisibleBy60(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
