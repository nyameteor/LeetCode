package main

import (
	"fmt"
	"slices"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	helper(root, 0, &res)
	return res
}

func helper(root *TreeNode, depth int, res *[]int) {
	if root == nil {
		return
	}

	if len(*res) == depth {
		*res = append(*res, root.Val)
	}

	helper(root.Right, depth+1, res)
	helper(root.Left, depth+1, res)
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected []int
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: []int{1, 3, 4, 5},
		},
		{
			input:    nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		actual := rightSideView(tc.input)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("rightSideView(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
