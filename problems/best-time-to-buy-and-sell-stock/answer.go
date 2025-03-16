package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	bestProfit := 0
	minPrice := math.MaxInt32
	for _, price := range prices {
		minPrice = min(minPrice, price)
		bestProfit = max(bestProfit, price-minPrice)
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
			expected: 5,
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
