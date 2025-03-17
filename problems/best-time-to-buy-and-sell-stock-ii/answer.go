package main

import "fmt"

func maxProfit(prices []int) int {
	bestProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			bestProfit += prices[i] - prices[i-1]
		}
	}
	return bestProfit
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := maxProfit(tc.input)
		if actual != tc.expected {
			fmt.Printf("maxProfit(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
