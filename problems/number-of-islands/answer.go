package main

import (
	"fmt"
	"slices"
)

func numIslands(grid [][]byte) int {
	return dfsInPlaceVisit(grid)
}

func dfsWithVisitArray(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		if seen[i][j] {
			return
		}
		seen[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	res := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == '1' && !seen[i][j] {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

// Performs DFS and marks the visited cells directly in the grid
func dfsInPlaceVisit(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		// Mark the current cell as visited (turn land into water)
		grid[i][j] = '0'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	res := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

func main() {
	testCases := []struct {
		input    [][]byte
		expected int
	}{
		{
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			input: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}

	solutions := []struct {
		name string
		fn   func(grid [][]byte) int
	}{
		{
			name: "Default Solution",
			fn:   numIslands,
		},
		{
			name: "DFS (Separate Visited Array)",
			fn:   dfsWithVisitArray,
		},
		{
			name: "DFS (In-place Visited)",
			fn:   dfsInPlaceVisit,
		},
	}

	deepClone := func(grid [][]byte) [][]byte {
		cloned := make([][]byte, len(grid))
		for i := range grid {
			cloned[i] = slices.Clone(grid[i])
		}
		return cloned
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			input := deepClone(tc.input)
			actual := solution.fn(input)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
