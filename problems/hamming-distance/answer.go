package main

import "fmt"

func hammingDistance(x int, y int) int {
	z := x ^ y
	return countSetBits(z)
}

func countSetBits(x int) int {
	count := 0

	for x > 0 {
		x &= x - 1
		count++
	}

	return count
}

func main() {
	type Input struct {
		x, y int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				x: 1,
				y: 4,
			},
			expected: 2,
		},
		{
			input: &Input{
				x: 3,
				y: 1,
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := hammingDistance(tc.input.x, tc.input.y)
		if actual != tc.expected {
			fmt.Printf("hammingDistance(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
