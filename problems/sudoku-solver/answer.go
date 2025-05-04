package main

import (
	"fmt"
	"slices"
	"strings"
)

func solveSudoku(board [][]byte) {
	var rows, cols, boxes [9][9]bool

	for row := range board {
		for col := range board[row] {
			if board[row][col] != '.' {
				num := board[row][col] - '1'
				box := (row/3)*3 + col/3
				rows[row][num], cols[col][num], boxes[box][num] = true, true, true
			}
		}
	}

	var solve func(index int) bool
	solve = func(index int) bool {
		if index == 81 {
			return true
		}

		row, col := index/9, index%9
		box := (row/3)*3 + col/3

		if board[row][col] != '.' {
			return solve(index + 1)
		}

		for num := range 9 {
			if rows[row][num] || cols[col][num] || boxes[box][num] {
				continue
			}

			board[row][col] = byte(num + '1')
			rows[row][num], cols[col][num], boxes[box][num] = true, true, true

			if solve(index + 1) {
				return true
			}

			board[row][col] = '.'
			rows[row][num], cols[col][num], boxes[box][num] = false, false, false
		}

		return false
	}

	solve(0)
}

func cloneBoard(board [][]byte) [][]byte {
	cloned := make([][]byte, len(board))
	for row := range board {
		cloned[row] = slices.Clone(board[row])
	}
	return cloned
}

func formatBoard(board [][]byte) string {
	var sb strings.Builder
	for row := range board {
		for col := range board[row] {
			sb.WriteByte(board[row][col])
			if col < len(board[row])-1 {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func main() {
	testCases := []struct {
		input    [][]byte
		expected [][]byte
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
			expected: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
	}

	for _, tc := range testCases {
		actual := cloneBoard(tc.input)
		solveSudoku(actual)
		if !slices.EqualFunc(actual, tc.expected, slices.Equal) {
			fmt.Printf("Input:\n%s\nOutput:\n%s\nExpected:\n%s\n",
				formatBoard(tc.input), formatBoard(actual), formatBoard(tc.expected))
		}
	}
}
