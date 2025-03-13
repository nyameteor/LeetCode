package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && checkSymmetric(root1.Left, root2.Right) && checkSymmetric(root1.Right, root2.Left)
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3},
				},
			},
			expected: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := isSymmetric(tc.input)
		if actual != tc.expected {
			fmt.Printf("isSymmetric(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
