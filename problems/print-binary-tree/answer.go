package main

import (
	"fmt"
	"reflect"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return 1 + max(height(node.Left), height(node.Right))
	}

	m := height(root)
	n := (1 << m) - 1

	grid := make([][]string, m)
	for i := range grid {
		grid[i] = make([]string, n)
	}

	var fill func(node *TreeNode, row, colStart, colEnd int)
	fill = func(node *TreeNode, row, colStart, colEnd int) {
		if node == nil {
			return
		}
		mid := (colStart + colEnd) / 2
		grid[row][mid] = fmt.Sprint(node.Val)
		fill(node.Left, row+1, colStart, mid-1)
		fill(node.Right, row+1, mid+1, colEnd)
	}

	fill(root, 0, 0, n-1)

	return grid
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected [][]string
	}{
		{
			input: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			expected: [][]string{
				{"", "1", ""},
				{"2", "", ""},
			},
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: [][]string{
				{"", "", "", "1", "", "", ""},
				{"", "2", "", "", "", "3", ""},
				{"", "", "4", "", "", "", ""},
			},
		},
	}

	for _, tc := range testCases {
		actual := printTree(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("printTree(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
