package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	// Handle overflow case: math.MinInt32 / -1 is out of 32-bit range.
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	isNegative := (dividend < 0) != (divisor < 0)

	quotient := divideAbsVals(abs(int64(dividend)), abs(int64(divisor)))

	if isNegative {
		return -int(quotient)
	}
	return int(quotient)
}

// Divides positive dividend and divisor, using int64 for safety
func divideAbsVals(dividend int64, divisor int64) int64 {
	if dividend < divisor {
		return 0
	}

	quotient := int64(0)
	curDivisor := divisor
	shiftCount := 0

	// Left shift divisor until it's larger than dividend
	for dividend >= (curDivisor << 1) {
		curDivisor <<= 1
		shiftCount++
	}

	// Subtract shifted divisor and adjust quotient
	for dividend >= divisor {
		if dividend >= curDivisor {
			dividend -= curDivisor
			quotient += 1 << shiftCount
		}
		curDivisor >>= 1
		shiftCount--
	}

	return quotient
}

func abs(x int64) int64 {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	type Input struct {
		dividend int
		divisor  int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				dividend: 10,
				divisor:  3,
			},
			expected: 3,
		},
		{
			input: &Input{
				dividend: 7,
				divisor:  -3,
			},
			expected: -2,
		},
		{
			input: &Input{
				dividend: 0,
				divisor:  1,
			},
			expected: 0,
		},
		{
			input: &Input{
				dividend: 1,
				divisor:  1,
			},
			expected: 1,
		},
		{
			input: &Input{
				dividend: math.MinInt32,
				divisor:  -1073741824,
			},
			expected: 2,
		},
		{
			input: &Input{
				dividend: math.MaxInt32,
				divisor:  math.MinInt32,
			},
			expected: 0,
		},
		{
			input: &Input{
				dividend: math.MinInt32,
				divisor:  -1,
			},
			expected: math.MaxInt32,
		},
	}

	for _, tc := range testCases {
		actual := divide(tc.input.dividend, tc.input.divisor)
		if actual != tc.expected {
			fmt.Printf("divide(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
