package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	maxGain(root, &res)
	return res
}

func maxGain(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	left := max(0, maxGain(root.Left, maxSum))
	right := max(0, maxGain(root.Right, maxSum))

	*maxSum = max(*maxSum, root.Val+left+right)

	return root.Val + max(left, right)
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 6,
		},
		{
			input: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 42,
		},
		{
			input: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			expected: 3,
		},
		{
			input:    &TreeNode{Val: -3},
			expected: -3,
		},
	}

	for _, tc := range testCases {
		actual := maxPathSum(tc.input)
		if actual != tc.expected {
			fmt.Printf("maxPathSum(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
