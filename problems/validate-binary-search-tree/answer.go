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

func isValidBST(root *TreeNode) bool {
	var check func(node *TreeNode, minVal, maxVal int) bool
	check = func(node *TreeNode, minVal, maxVal int) bool {
		if node == nil {
			return true
		}
		if node.Val <= minVal || node.Val >= maxVal {
			return false
		}
		return check(node.Left, minVal, node.Val) && check(node.Right, node.Val, maxVal)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			input: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			expected: false,
		},
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 6},
				},
			},
			expected: true,
		},
		{
			input: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: false,
		},
		{
			input: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			input:    &TreeNode{Val: math.MaxInt32},
			expected: true,
		},
	}

	for _, tc := range testCases {
		actual := isValidBST(tc.input)
		if actual != tc.expected {
			fmt.Printf("isValidBST(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
