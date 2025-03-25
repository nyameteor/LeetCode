package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	row, col := 0, n-1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
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
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			expected: true,
		},
		{
			input: &Input{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 20,
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
