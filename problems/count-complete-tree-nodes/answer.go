package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)

	if left == right {
		return 1<<left + countNodes(root.Right)
	} else {
		return 1<<right + countNodes(root.Left)
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + height(root.Left)
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
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: 6,
		},
		{
			input:    nil,
			expected: 0,
		},
		{
			input:    &TreeNode{Val: 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := countNodes(tc.input)
		if actual != tc.expected {
			fmt.Printf("countNodes(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
