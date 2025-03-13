package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / fastPow(x, int64(-n))
	}
	return fastPow(x, int64(n))
}

// Computes power for positive exponents using exponentiation by squaring
// Use int64 to handle large values safely (e.g., -math.MinInt32)
func fastPow(x float64, n int64) float64 {
	if n == 0 {
		return 1
	}

	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return x * half * half
}

func main() {
	type Input struct {
		x float64
		n int
	}

	testCases := []struct {
		input    *Input
		expected float64
	}{
		{
			input: &Input{
				x: 2.00000,
				n: 10,
			},
			expected: 1024.00000,
		},
		{
			input: &Input{
				x: 2.10000,
				n: 3,
			},
			expected: 9.26100,
		},
		{
			input: &Input{
				x: 2.00000,
				n: -2,
			},
			expected: 0.25000,
		},
		{
			input: &Input{
				x: 1.00000,
				n: math.MinInt32,
			},
			expected: 1.00000,
		},
		{
			input: &Input{
				x: 2.00000,
				n: math.MinInt32,
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := myPow(tc.input.x, tc.input.n)
		if math.Nextafter(actual, tc.expected) != tc.expected {
			fmt.Printf("myPow(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
