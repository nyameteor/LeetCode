package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	var postions [][]int
	var res [][]string
	for j := 0; j < n; j++ {
		solver(&res, &postions, n, 0, j)
	}

	return res
}

func solver(res *[][]string, positions *[][]int, n int, i int, j int) {
	if i >= n {
		return
	}

	if !isPosValid(positions, i, j) {
		return
	}

	(*positions) = append((*positions), []int{i, j})

	if len(*positions) == n {
		(*res) = append((*res), convertRes(positions, n))
	}

	for k := 0; k < n; k++ {
		solver(res, positions, n, i+1, k)
	}

	(*positions) = (*positions)[:len(*positions)-1]
}

func isPosValid(positions *[][]int, i int, j int) bool {
	for k := range *positions {
		x := (*positions)[k][0]
		y := (*positions)[k][1]

		if i == x || j == y || i-x == j-y || i-x == -(j-y) {
			return false
		}
	}

	return true
}

func convertRes(positions *[][]int, n int) []string {
	res := make([]string, n)
	for i := range *positions {
		x := (*positions)[i][0]
		y := (*positions)[i][1]

		s := ""
		for j := 0; j < n; j++ {
			if j == y {
				s += "Q"
			} else {
				s += "."
			}
		}

		res[x] = s
	}

	return res
}

func printResult(res [][]string) {
	for i := range res {
		for j := range res[i] {
			fmt.Print(res[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	printResult(solveNQueens(1))
	printResult(solveNQueens(4))
	printResult(solveNQueens(8))
}
