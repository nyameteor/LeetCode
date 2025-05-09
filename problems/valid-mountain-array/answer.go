package main

import "fmt"

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	// Walk up
	i := 0
	for ; i < n-1 && arr[i] < arr[i+1]; i++ {
	}

	// Peak can't be first or last
	if i == 0 || i == n-1 {
		return false
	}

	// Walk down
	for ; i < n-1 && arr[i] > arr[i+1]; i++ {
	}

	return i == n-1
}

func main() {
	testCases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{2, 1},
			expected: false,
		},
		{
			input:    []int{3, 5, 5},
			expected: false,
		},
		{
			input:    []int{0, 3, 2, 1},
			expected: true,
		},
		{
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: false,
		},
		{
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := validMountainArray(tc.input)
		if actual != tc.expected {
			fmt.Printf("validMountainArray(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
