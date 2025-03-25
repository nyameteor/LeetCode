package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	numRows, numCols := len(matrix), len(matrix[0])

	l, r := 0, numRows*numCols-1
	for l <= r {
		m := l + (r-l)/2

		row, col := m/numCols, m%numCols
		val := matrix[row][col]

		if val == target {
			return true
		} else if val > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

func main() {
	type Input struct {
		matrix [][]int
		target int
	}

	testCases := []struct {
		input    *Input
		expected bool
	}{
		{
			input: &Input{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 3,
			},
			expected: true,
		},
		{
			input: &Input{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 13,
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := searchMatrix(tc.input.matrix, tc.input.target)
		if actual != tc.expected {
			fmt.Printf("searchMatrix(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
