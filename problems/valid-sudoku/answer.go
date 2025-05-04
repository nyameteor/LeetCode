package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}

	for row := range 9 {
		for col := range 9 {
			if board[row][col] == '.' {
				continue
			}

			num := board[row][col] - '1'
			box := (row/3)*3 + col/3

			if rows[row][num] || cols[col][num] || boxes[box][num] {
				return false
			}

			rows[row][num], cols[col][num], boxes[box][num] = true, true, true
		}
	}

	return true
}

func main() {
	testCases := []struct {
		input    [][]byte
		expected bool
	}{
		{
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: true,
		},
		{
			input: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := isValidSudoku(tc.input)
		if actual != tc.expected {
			fmt.Printf("isValidSudoku(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
