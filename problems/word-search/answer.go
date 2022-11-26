package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res := false
			dfs(&res, board, seen, i, j, m, n, word, 0)
			if res {
				return true
			}
		}
	}

	return false
}

func dfs(res *bool, board [][]byte, seen [][]bool, i, j, m, n int, word string, pos int) {
	if pos == len(word) {
		*res = true
		return
	}

	if !(i >= 0 && i < m && j >= 0 && j < n) {
		return
	}

	if seen[i][j] || board[i][j] != word[pos] {
		return
	}

	seen[i][j] = true
	pos++
	dfs(res, board, seen, i-1, j, m, n, word, pos)
	dfs(res, board, seen, i+1, j, m, n, word, pos)
	dfs(res, board, seen, i, j-1, m, n, word, pos)
	dfs(res, board, seen, i, j+1, m, n, word, pos)
	pos--
	seen[i][j] = false
}

func main() {
	var board [][]byte
	var word string

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "ABCCED"
	fmt.Println(exist(board, word) == true)

	word = "SEE"
	fmt.Println(exist(board, word) == true)

	word = "ABCB"
	fmt.Println(exist(board, word) == false)

	board = [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}

	word = "AAB"
	fmt.Println(exist(board, word) == true)

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "ABCESEEEFS"
	fmt.Println(exist(board, word) == true)

	board = [][]byte{
		{'A'},
	}

	word = "A"
	fmt.Println(exist(board, word) == true)
}
