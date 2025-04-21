package main

import "fmt"

func arrangeCoins(n int) int {
	left, right := 0, n

	for left <= right {
		mid := left + (right-left)/2
		sum := mid * (mid + 1) / 2

		if sum == n {
			return mid
		} else if sum < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    5,
			expected: 2,
		},
		{
			input:    8,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		actual := arrangeCoins(tc.input)
		if actual != tc.expected {
			fmt.Printf("arrangeCoins(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
