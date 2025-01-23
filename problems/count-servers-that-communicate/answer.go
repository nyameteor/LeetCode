package main

import "fmt"

func countServers(grid [][]int) int {
	numRows := len(grid)
	numCols := len(grid[0])

	rowCounts := make([]int, numCols)
	colCounts := make([]int, numRows)

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if grid[row][col] == 1 {
				rowCounts[col]++
				colCounts[row]++
			}
		}
	}

	res := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if grid[row][col] == 1 {
				if rowCounts[col] > 1 || colCounts[row] > 1 {
					res++
				}
			}
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{1, 0},
				{0, 1},
			},
			expected: 0,
		},
		{
			input: [][]int{
				{1, 0},
				{1, 1},
			},
			expected: 3,
		},
		{
			input: [][]int{
				{1, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		actual := countServers(tc.input)
		if actual != tc.expected {
			fmt.Printf("countServers(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
