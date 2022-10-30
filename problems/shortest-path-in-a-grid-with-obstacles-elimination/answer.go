package main

import "fmt"

type cell struct {
	i int
	j int
	k int
}

type pos struct {
	i int
	j int
}

func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	// edge case
	if k >= m+n-2 {
		return m + n - 2
	}

	q := []cell{{0, 0, k}}
	level := 0
	cur := cell{}

	// keep track of visited cells
	seen := map[cell]struct{}{}

	// directions of next step
	dirs := []pos{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur, q = q[0], q[1:]

			// reach the end
			if cur.i == m-1 && cur.j == n-1 {
				return level
			}

			for _, dir := range dirs {
				ii := cur.i + dir.i
				jj := cur.j + dir.j
				kk := cur.k
				// is cell position valid
				if !(0 <= ii && ii <= m-1 && 0 <= jj && jj <= n-1) {
					continue
				}
				// try to eliminate obstacle
				if grid[ii][jj] == 1 {
					if kk <= 0 {
						continue
					}
					kk -= 1
				}
				c := cell{ii, jj, kk}

				if _, ok := seen[c]; ok {
					continue
				}
				// remember the new cell
				seen[c] = struct{}{}
				q = append(q, c)
			}
		}
		level++
	}

	return -1
}

func main() {
	var grid [][]int
	var k int
	var res int

	grid = [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}
	k = 1
	res = 6
	fmt.Println(shortestPath(grid, k) == res)

	grid = [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
	}
	k = 1
	res = 9
	fmt.Println(shortestPath(grid, k) == res)

	grid = [][]int{
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{0, 0, 0, 1, 0, 0, 1, 0},
	}
	k = 1
	res = 13
	fmt.Println(shortestPath(grid, k) == res)

	grid = [][]int{
		{0, 1, 1},
		{1, 1, 1},
		{1, 0, 0},
	}
	k = 1
	res = -1
	fmt.Println(shortestPath(grid, k) == res)
}
