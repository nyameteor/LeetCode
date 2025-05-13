package main

import "fmt"

func orangesRotting(grid [][]int) int {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	queue := [][2]int{}
	level := 0
	fresh := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	for len(queue) > 0 {
		n := len(queue)
		for range n {
			cur := queue[0]
			queue = queue[1:]

			for _, direction := range directions {
				r, c := cur[0]+direction[0], cur[1]+direction[1]
				if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 {
					grid[r][c] = 2
					fresh--
					queue = append(queue, [2]int{r, c})
				}
			}
		}
		level++
	}

	if fresh > 0 {
		return -1
	}

	return level - 1
}

func main() {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expected: 4,
		},
		{
			input:    [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			expected: -1,
		},
		{
			input:    [][]int{{0, 2}},
			expected: 0,
		},
		{
			input:    [][]int{{0}},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := orangesRotting(tc.input)
		if actual != tc.expected {
			fmt.Printf("orangesRotting(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
