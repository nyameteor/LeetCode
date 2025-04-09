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

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	even := true

	for len(queue) > 0 {
		n := len(queue)
		vals := make([]int, n)

		for i := range n {
			node := queue[0]
			queue = queue[1:]

			if even {
				vals[i] = node.Val
			} else {
				vals[n-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, vals)
		even = !even
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
			expected: [][]int{{3}, {20, 9}, {15, 7}},
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
		actual := zigzagLevelOrder(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("zigzagLevelOrder(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
