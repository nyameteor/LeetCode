package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	i := 0
	j := 0
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	cur_d := 0
	for cur := head; cur != nil; cur = cur.Next {
		res[i][j] = cur.Val
		ni := i + directions[cur_d][0]
		nj := j + directions[cur_d][1]

		if ni < 0 || ni >= m || nj < 0 || nj >= n || res[ni][nj] != -1 {
			cur_d = (cur_d + 1) % 4
		}
		i += directions[cur_d][0]
		j += directions[cur_d][1]
	}

	return res
}

func main() {
	var m int
	var n int
	var head *ListNode
	var res [][]int

	m = 3
	n = 5
	head = &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8,
		&ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2,
			&ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}
	res = spiralMatrix(m, n, head)
	fmt.Printf("%v\n", res)

	m = 1
	n = 4
	head = &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	res = spiralMatrix(m, n, head)
	fmt.Printf("%v\n", res)
}
