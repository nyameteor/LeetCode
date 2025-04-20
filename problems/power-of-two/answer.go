package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return checkWithBitwise(n)
}

func checkWithBitCount(n int) bool {
	if n <= 0 {
		return false
	}
	for n&1 == 0 {
		n >>= 1
	}
	return n == 1
}

func checkWithBitwise(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func main() {
	testCases := []struct {
		input    int
		expected bool
	}{
		{
			input:    1,
			expected: true,
		},
		{
			input:    16,
			expected: true,
		},
		{
			input:    3,
			expected: false,
		},
		{
			input:    0,
			expected: false,
		},
	}

	solutions := []struct {
		name string
		fn   func(int) bool
	}{
		{
			name: "Default Solution",
			fn:   isPowerOfTwo,
		},
		{
			name: "Bit Count Method",
			fn:   checkWithBitCount,
		},
		{
			name: "Bitwise Method",
			fn:   checkWithBitwise,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
