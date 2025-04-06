package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	type Input struct {
		root, p, q *TreeNode
	}

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}

	testCases := []struct {
		input    *Input
		expected *TreeNode
	}{
		{
			input: &Input{
				root: root,
				p:    root.Left,
				q:    root.Right,
			},
			expected: root,
		},
		{
			input: &Input{
				root: root,
				p:    root.Left,
				q:    root.Left.Right.Right,
			},
			expected: root.Left,
		},
	}

	for _, tc := range testCases {
		actual := lowestCommonAncestor(tc.input.root, tc.input.p, tc.input.q)
		if actual != tc.expected {
			fmt.Printf("lowestCommonAncestor(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
