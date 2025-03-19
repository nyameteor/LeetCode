package main

import (
	"fmt"
	"reflect"
	"slices"
)

func rotate(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
}

func transpose(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverse(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for l, r := 0, len(matrix)-1; l < r; l, r = l+1, r-1 {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
		}
	}
}

func main() {
	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, tc := range testCases {
		inputCopy := make([][]int, len(tc.input))
		for i := range tc.input {
			inputCopy[i] = slices.Clone(tc.input[i])
		}

		rotate(tc.input)
		actual := tc.input
		tc.input = inputCopy

		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("rotate(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
