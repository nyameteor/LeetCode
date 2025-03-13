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

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		n := len(q)
		row := []int{}

		for range n {
			cur := q[0]
			q = q[1:]

			row = append(row, cur.Val)

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}

		res = append(res, row)
	}

	return res
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected [][]int
	}{
		{
			input: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			input:    &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			input:    nil,
			expected: [][]int{},
		},
	}

	for _, tc := range testCases {
		actual := levelOrder(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("levelOrder(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
