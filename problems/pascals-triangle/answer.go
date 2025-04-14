package main

import (
	"fmt"
	"reflect"
)

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    int
		expected [][]int
	}{
		{
			input:    5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			input:    1,
			expected: [][]int{{1}},
		},
	}

	for _, tc := range testCases {
		actual := generate(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("generate(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
