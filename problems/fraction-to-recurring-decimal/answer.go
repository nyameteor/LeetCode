package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var res strings.Builder

	if (numerator < 0) != (denominator < 0) {
		res.WriteByte('-')
	}

	numerator, denominator = abs(numerator), abs(denominator)

	// Integer part
	res.WriteString(strconv.Itoa(numerator / denominator))
	numerator %= denominator
	if numerator == 0 {
		return res.String()
	}

	res.WriteByte('.')

	// Track position of remainders
	remainderPos := map[int]int{}

	// Fraction part
	for numerator != 0 {
		if pos, ok := remainderPos[numerator]; ok {
			s := res.String()
			return s[:pos] + "(" + s[pos:] + ")"
		}

		remainderPos[numerator] = res.Len()
		numerator *= 10
		res.WriteByte('0' + byte(numerator/denominator))
		numerator %= denominator
	}

	return res.String()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	type Input struct {
		numerator, denominator int
	}

	testCases := []struct {
		input    *Input
		expected string
	}{
		{
			input:    &Input{1, 2},
			expected: "0.5",
		},
		{
			input:    &Input{2, 1},
			expected: "2",
		},
		{
			input:    &Input{4, 333},
			expected: "0.(012)",
		},
		{
			input:    &Input{22, 7},
			expected: "3.(142857)",
		},
		{
			input:    &Input{0, -5},
			expected: "0",
		},
	}

	for _, tc := range testCases {
		actual := fractionToDecimal(tc.input.numerator, tc.input.denominator)
		if actual != tc.expected {
			fmt.Printf("fractionToDecimal(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
