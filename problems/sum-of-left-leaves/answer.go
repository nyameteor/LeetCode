package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(node *TreeNode, isLeft bool) int
	dfs = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil && isLeft {
			return node.Val
		}
		return dfs(node.Left, true) + dfs(node.Right, false)
	}

	return dfs(root, false)
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected int
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
			expected: 24,
		},
		{
			input:    &TreeNode{Val: 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := sumOfLeftLeaves(tc.input)
		if actual != tc.expected {
			fmt.Printf("sumOfLeftLeaves(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
