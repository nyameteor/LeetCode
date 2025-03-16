package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := invertTree(root.Left)
	r := invertTree(root.Right)

	root.Left, root.Right = r, l

	return root
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected *TreeNode
	}{
		{
			input: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
	}

	var treesEqual func(root1 *TreeNode, root2 *TreeNode) bool
	treesEqual = func(root1 *TreeNode, root2 *TreeNode) bool {
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil {
			return false
		}
		return root1.Val == root2.Val && treesEqual(root1.Left, root2.Left) && treesEqual(root1.Right, root2.Right)
	}

	for _, tc := range testCases {
		actual := invertTree(tc.input)
		if !treesEqual(actual, tc.expected) {
			fmt.Printf("invertTree(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
