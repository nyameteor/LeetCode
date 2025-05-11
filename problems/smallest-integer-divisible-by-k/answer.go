package main

import "fmt"

func smallestRepunitDivByK(k int) int {
	return optimized(k)
}

func general(k int) int {
	seen := make(map[int]bool)
	remainder := 1 % k

	for length := 1; ; length++ {
		if remainder == 0 {
			return length
		}
		if seen[remainder] {
			break
		}
		seen[remainder] = true
		remainder = (10*remainder + 1) % k
	}

	return -1
}

func optimized(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	reminder := 1 % k

	for length := 1; length <= k; length++ {
		if reminder == 0 {
			return length
		}
		reminder = (10*reminder + 1) % k
	}

	return -1
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: -1,
		},
		{
			input:    3,
			expected: 3,
		},
	}

	solutions := []struct {
		name string
		fn   func(int) int
	}{
		{
			name: "Default Solution",
			fn:   smallestRepunitDivByK,
		},
		{
			name: "General Solution",
			fn:   general,
		},
		{
			name: "Optimized Solution",
			fn:   optimized,
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
