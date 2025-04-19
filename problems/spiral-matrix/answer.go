package main

import (
	"fmt"
	"slices"
)

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	top, left := 0, 0
	bottom, right := m-1, n-1

	res := make([]int, 0, m*n)
	for top <= bottom && left <= right {
		// left to right
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		top++

		// top to bottom
		for r := top; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		right--

		// right to left, if there's still a row left
		if top <= bottom {
			for c := right; c >= left; c-- {
				res = append(res, matrix[bottom][c])
			}
			bottom--
		}

		// bottom to top, if there's still a column left
		if left <= right {
			for r := bottom; r >= top; r-- {
				res = append(res, matrix[r][left])
			}
			left++
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			input:    [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, tc := range testCases {
		actual := spiralOrder(tc.input)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("spiralOrder(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
