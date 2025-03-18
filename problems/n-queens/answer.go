package main

import "fmt"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	dfs(make([]int, 0), 0, n, &res)
	return res
}

func dfs(positions []int, row int, n int, res *[][]string) {
	if row == n {
		*res = append(*res, buildBoard(positions))
		return
	}

	for col := range n {
		if isValidPosition(positions, col) {
			dfs(append(positions, col), row+1, n, res)
		}
	}
}

func isValidPosition(positions []int, newColumn int) bool {
	newRow := len(positions)
	for row, col := range positions {
		if col == newColumn || row-newRow == col-newColumn || row-newRow == -(col-newColumn) {
			return false
		}
	}
	return true
}

func buildBoard(positions []int) []string {
	n := len(positions)
	board := make([]string, n)
	for r, c := range positions {
		row := make([]rune, n)
		for j := range row {
			row[j] = '.'
		}
		row[c] = 'Q'
		board[r] = string(row)
	}
	return board
}

func main() {
	fmt.Println(solveNQueens(1))
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(8))
}
