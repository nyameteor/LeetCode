package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	memo := make([][]int, m)
	for i := range m {
		memo[i] = make([]int, n)
		for j := range n {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if matrix[i][j] == '0' {
			return 0
		}

		memo[i][j] = min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
		return memo[i][j]
	}

	maxSide := 0
	for i := range m {
		for j := range n {
			maxSide = max(maxSide, dfs(i, j))
		}
	}

	return maxSide * maxSide
}

func main() {
	testCases := []struct {
		input    [][]byte
		expected int
	}{
		{
			input: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 4,
		},
		{
			input: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			expected: 1,
		},
		{
			input: [][]byte{
				{'0'},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := maximalSquare(tc.input)
		if actual != tc.expected {
			fmt.Printf("maximalSquare(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
