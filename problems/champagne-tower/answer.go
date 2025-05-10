package main

import (
	"fmt"
	"math"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	curRows := make([]float64, 1)
	curRows[0] = float64(poured)

	for i := range query_row {
		nextRows := make([]float64, i+2)
		for j := range curRows {
			if curRows[j] < 1 {
				continue
			}
			overflow := curRows[j] - 1
			nextRows[j] += overflow / 2
			nextRows[j+1] += overflow / 2
		}
		curRows = nextRows
	}

	if curRows[query_glass] > 1 {
		return 1
	}
	return float64(curRows[query_glass])
}

func main() {
	type Input struct {
		poured, query_row, query_glass int
	}

	testCases := []struct {
		input    *Input
		expected float64
	}{
		{
			input: &Input{
				poured:      1,
				query_row:   1,
				query_glass: 1,
			},
			expected: 0,
		},
		{
			input: &Input{
				poured:      2,
				query_row:   1,
				query_glass: 1,
			},
			expected: 0.5,
		},
		{
			input: &Input{
				poured:      100000009,
				query_row:   33,
				query_glass: 17,
			},
			expected: 1,
		},
	}

	floatEqual := func(a, b, epsilon float64) bool {
		return math.Abs(a-b) <= epsilon
	}
	epsilon := 1e-9

	for _, tc := range testCases {
		actual := champagneTower(tc.input.poured, tc.input.query_row, tc.input.query_glass)
		if !floatEqual(actual, tc.expected, epsilon) {
			fmt.Printf("champagneTower(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
