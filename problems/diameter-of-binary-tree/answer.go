package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	lh := dfs(root.Left, res)
	rh := dfs(root.Right, res)

	*res = max(*res, lh+rh)

	return 1 + max(lh, rh)
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: 3,
		},
		{
			input: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := diameterOfBinaryTree(tc.input)
		if actual != tc.expected {
			fmt.Printf("diameterOfBinaryTree(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
