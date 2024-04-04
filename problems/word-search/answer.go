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
			if dfs(board, seen, i, j, m, n, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, seen [][]bool, i, j, m, n int, word string, pos int) bool {
	if pos == len(word) {
		return true
	}

	if !(0 <= i && i < m) || !(0 <= j && j < n) {
		return false
	}

	if seen[i][j] || board[i][j] != word[pos] {
		return false
	}

	seen[i][j] = true
	pos++
	if dfs(board, seen, i-1, j, m, n, word, pos) {
		return true
	}
	if dfs(board, seen, i+1, j, m, n, word, pos) {
		return true
	}
	if dfs(board, seen, i, j-1, m, n, word, pos) {
		return true
	}
	if dfs(board, seen, i, j+1, m, n, word, pos) {
		return true
	}
	pos--
	seen[i][j] = false
	return false
}

func main() {
	var board [][]byte
	var word string
	var res bool

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "ABCCED"
	res = exist(board, word)
	fmt.Println(res == true)

	word = "SEE"
	res = exist(board, word)
	fmt.Println(res == true)

	word = "ABCB"
	res = exist(board, word)
	fmt.Println(res == false)

	board = [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}

	word = "AAB"
	res = exist(board, word)
	fmt.Println(res == true)

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "ABCESEEEFS"
	res = exist(board, word)
	fmt.Println(res == true)

	board = [][]byte{
		{'A'},
	}

	word = "A"
	res = exist(board, word)
	fmt.Println(res == true)
}
