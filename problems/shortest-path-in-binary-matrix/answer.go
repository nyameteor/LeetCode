package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}

	queue := [][2]int{{0, 0}}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[0][0] = true

	level := 0
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for len(queue) > 0 {
		levelLen := len(queue)
		for range levelLen {
			cur := queue[0]
			queue = queue[1:]

			if cur[0] == n-1 && cur[1] == n-1 {
				return level + 1
			}

			for _, direction := range directions {
				r, c := cur[0]+direction[0], cur[1]+direction[1]
				if r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0 && !seen[r][c] {
					queue = append(queue, [2]int{r, c})
					seen[r][c] = true
				}
			}
		}
		level++
	}

	return -1
}

func main() {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{0, 1}, {1, 0}},
			expected: 2,
		},
		{
			input:    [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			expected: 4,
		},
		{
			input:    [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		actual := shortestPathBinaryMatrix(tc.input)
		if actual != tc.expected {
			fmt.Printf("shortestPathBinaryMatrix(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
