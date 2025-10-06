package main

import (
	"fmt"
	"reflect"
)

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	res := make([]int, m*n)

	up := true
	i, j := 0, 0

	for k := 0; k < m*n; k++ {
		res[k] = mat[i][j]

		if up {
			// moving up-right
			switch {
			case j == n-1:
				// hit right edge, move down and reverse direction
				i++
				up = false
			case i == 0:
				// hit top edge, move right and reverse direction
				j++
				up = false
			default:
				i--
				j++
			}
		} else {
			// moving down-left
			switch {
			case i == m-1:
				// hit bottom edge, move right and reverse direction
				j++
				up = true
			case j == 0:
				// hit left edge, move down and reverse direction
				i++
				up = true
			default:
				i++
				j--
			}
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
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		actual := findDiagonalOrder(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("findDiagonalOrder(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
